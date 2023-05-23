package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"fmt"
	"sort"
	"time"

	"golang.org/x/xerrors"
)

func (f *SubFunction) BuyerSellerDetection(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.BuyerSellerDetection, error) {
	var buyerOrSeller string
	metaData := psdc.MetaData

	if *metaData.BusinessPartnerID == *sdc.Orders.Buyer && *metaData.BusinessPartnerID != *sdc.Orders.Seller {
		buyerOrSeller = "Buyer"
	} else if *metaData.BusinessPartnerID != *sdc.Orders.Buyer && *metaData.BusinessPartnerID == *sdc.Orders.Seller {
		buyerOrSeller = "Seller"
	} else {
		return nil, fmt.Errorf("business_partnerがBuyerまたはSellerと一致しません")
	}

	buyerSellerDetection := psdc.ConvertToBuyerSellerDetection(sdc, buyerOrSeller)

	return buyerSellerDetection, nil
}

func (f *SubFunction) HeaderBPCustomerSupplier(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.HeaderBPCustomerSupplier, error) {
	var rows *sql.Rows
	var err error

	buyerSellerDetection := psdc.BuyerSellerDetection
	if buyerSellerDetection.BuyerOrSeller == "Seller" {
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Customer, Currency, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_data
		WHERE (BusinessPartner, Customer) = (?, ?);`, buyerSellerDetection.BusinessPartnerID, buyerSellerDetection.Buyer,
		)
		if err != nil {
			return nil, err
		}
		psdc.HeaderBPCustomer, err = psdc.ConvertToHeaderBPCustomer(rows)
		if err != nil {
			return nil, err
		}

		rows, err = f.db.Query(
			`SELECT BusinessPartner, Supplier, Currency, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data
		WHERE (BusinessPartner, Supplier) = (?, ?);`, buyerSellerDetection.Buyer, buyerSellerDetection.Seller,
		)
		if err != nil {
			return nil, err
		}
		psdc.HeaderBPSupplier, err = psdc.ConvertToHeaderBPSupplier(rows)
		if err != nil {
			return nil, err
		}
	} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Supplier, Currency, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data
		WHERE (BusinessPartner, Supplier) = (?, ?);`, buyerSellerDetection.BusinessPartnerID, buyerSellerDetection.Seller,
		)
		if err != nil {
			return nil, err
		}
		psdc.HeaderBPSupplier, err = psdc.ConvertToHeaderBPSupplier(rows)
		if err != nil {
			return nil, err
		}

		rows, err = f.db.Query(
			`SELECT BusinessPartner, Customer, Currency, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_data
		WHERE (BusinessPartner, Customer) = (?, ?);`, buyerSellerDetection.Seller, buyerSellerDetection.Buyer,
		)
		if err != nil {
			return nil, err
		}
		psdc.HeaderBPCustomer, err = psdc.ConvertToHeaderBPCustomer(rows)
		if err != nil {
			return nil, err
		}
	}

	data := psdc.ConvertToHeaderBPCustomerSupplier()

	return data, err
}

func (f *SubFunction) CalculateOrderID(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateOrderID, error) {
	metaData := psdc.MetaData
	dataKey := psdc.ConvertToCalculateOrderIDKey()

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateOrderIDQueryGets(rows)
	if err != nil {
		return nil, err
	}

	orderIDLatestNumber := dataQueryGets.OrderIDLatestNumber
	orderID := *dataQueryGets.OrderIDLatestNumber + 1

	data := psdc.ConvertToCalculateOrderID(orderIDLatestNumber, orderID)

	return data, err
}

func (f *SubFunction) InvoiceDocumentDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.InvoiceDocumentDate, error) {

	dataKey := psdc.ConvertToPaymentTermsKey()

	dataKey.PaymentTerms = psdc.HeaderBPCustomerSupplier.PaymentTerms

	rows, err := f.db.Query(
		`SELECT PaymentTerms, BaseDate, BaseDateCalcAddMonth, BaseDateCalcFixedDate, PaymentDueDateCalcAddMonth, PaymentDueDateCalcFixedDate
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data
			WHERE PaymentTerms = ?;`, dataKey.PaymentTerms,
	)
	if err != nil {
		return nil, err
	}

	psdc.PaymentTerms, err = psdc.ConvertToPaymentTerms(rows)
	if err != nil {
		return nil, err
	}

	if sdc.Orders.InvoiceDocumentDate != nil {
		if *sdc.Orders.InvoiceDocumentDate != "" {
			data := psdc.ConvertToInvoiceDocumentDate(sdc)
			return data, nil
		}
	}

	requestedDeliveryDate := psdc.ConvertToRequestedDeliveryDate(sdc)

	calculateInvoiceDocumentDate, err := CalculateInvoiceDocumentDate(psdc, requestedDeliveryDate.RequestedDeliveryDate, psdc.PaymentTerms)
	if err != nil {
		return nil, err
	}

	data := psdc.ConvertToCaluculateInvoiceDocumentDate(sdc, calculateInvoiceDocumentDate)

	return data, err
}

func CalculateInvoiceDocumentDate(
	psdc *api_processing_data_formatter.SDC,
	requestedDeliveryDate string,
	paymentTerms *[]api_processing_data_formatter.PaymentTerms,
) (string, error) {

	format := "2006-01-02"
	t, err := time.Parse(format, requestedDeliveryDate)
	if err != nil {
		return "", err
	}

	sort.Slice(*paymentTerms, func(i, j int) bool {
		return (*paymentTerms)[i].BaseDate < (*paymentTerms)[j].BaseDate
	})

	day := t.Day()
	for i, v := range *paymentTerms {
		if day <= v.BaseDate {
			t = time.Date(t.Year(), t.Month()+time.Month(*v.BaseDateCalcAddMonth)+1, 0, 0, 0, 0, 0, time.UTC)
			if *v.BaseDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.BaseDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
			break
		}
		if i == len(*paymentTerms)-1 {
			return "", xerrors.Errorf("'data_platform_payment_terms_payment_terms_data'テーブルが不適切です。")
		}
	}

	res := t.Format(format)

	return res, nil
}

func (f *SubFunction) PaymentDueDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.PaymentDueDate, error) {

	if sdc.Orders.PaymentDueDate != nil {
		if *sdc.Orders.PaymentDueDate != "" {
			data := psdc.ConvertToPaymentDueDate(sdc)
			return data, nil
		}
	}

	calculatePaymentDueDate, err := CalculatePaymentDueDate(psdc.InvoiceDocumentDate.InvoiceDocumentDate, psdc.PaymentTerms)
	if err != nil {
		return nil, err
	}

	data := psdc.ConvertToCaluculatePaymentDueDate(calculatePaymentDueDate)

	return data, err
}

func CalculatePaymentDueDate(
	invoiceDocumentDate string,
	paymentTerms *[]api_processing_data_formatter.PaymentTerms,
) (string, error) {

	format := "2006-01-02"
	t, err := time.Parse(format, invoiceDocumentDate)
	if err != nil {
		return "", err
	}

	sort.Slice(*paymentTerms, func(i, j int) bool {
		return (*paymentTerms)[i].BaseDate < (*paymentTerms)[j].BaseDate
	})

	day := t.Day()
	if day == time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day() {
		day = 31
	}
	for i, v := range *paymentTerms {
		if day <= *v.BaseDateCalcFixedDate {
			t = time.Date(t.Year(), t.Month()+time.Month(*v.PaymentDueDateCalcAddMonth)+1, 0, 0, 0, 0, 0, time.UTC)
			if *v.PaymentDueDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.PaymentDueDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
			break
		}
		if i == len(*paymentTerms)-1 {
			t = time.Date(t.Year(), t.Month()+time.Month(*v.PaymentDueDateCalcAddMonth)+2, 0, 0, 0, 0, 0, time.UTC)
			if *v.PaymentDueDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.PaymentDueDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
		}
	}

	res := t.Format(format)

	return res, nil
}

func (f *SubFunction) NetPaymentDays(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.NetPaymentDays, error) {

	if sdc.Orders.NetPaymentDays != nil {
		data := psdc.ConvertToNetPaymentDays(sdc)
		return data, nil
	}

	calculateNetPaymentDays, err := CalculateNetPaymentDays(psdc.InvoiceDocumentDate.InvoiceDocumentDate, psdc.PaymentDueDate.PaymentDueDate)
	if err != nil {
		return nil, err
	}

	data := psdc.ConvertToCaluculateNetPaymentDays(calculateNetPaymentDays)

	return data, err
}

func CalculateNetPaymentDays(
	invoiceDocumentDate string,
	paymentDueDate string,
) (int, error) {

	format := "2006-01-02"
	tb, err := time.Parse(format, invoiceDocumentDate)
	if err != nil {
		return 0, err
	}

	tp, err := time.Parse(format, paymentDueDate)
	if err != nil {
		return 0, err
	}

	res := int(tp.Sub(tb).Hours() / 24)

	return res, nil
}

func (f *SubFunction) OverallDocReferenceStatus(
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.OverallDocReferenceStatus, error) {
	var overallDocReferenceStatus string
	serviceLabel := psdc.OrderReferenceType.ServiceLabel

	if serviceLabel == "QUOTATIONS" {
		overallDocReferenceStatus = "QT"
	} else if serviceLabel == "INQUIRIES" {
		overallDocReferenceStatus = "IN"
	} else if serviceLabel == "PURCHASE_REQUISITION" {
		overallDocReferenceStatus = "PR"
	}

	data := psdc.ConvertToOverallDocReferenceStatus(overallDocReferenceStatus)

	return data, nil
}

func (f *SubFunction) PricingDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.PricingDate {
	var data *api_processing_data_formatter.PricingDate

	if sdc.Orders.PricingDate != nil {
		if *sdc.Orders.PricingDate != "" {
			data = psdc.ConvertToPricingDate(*sdc.Orders.PricingDate)
		}
	} else {
		data = psdc.ConvertToPricingDate(GetDateStr())
	}

	return data
}

func (f *SubFunction) PriceDetnExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.PriceDetnExchangeRate, error) {
	data := psdc.ConvertToPriceDetnExchangeRate(sdc)

	return data, nil
}

func (f *SubFunction) AccountingExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.AccountingExchangeRate, error) {
	data := psdc.ConvertToAccountingExchangeRate(sdc)

	return data, nil
}

func (f *SubFunction) Currency(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) error {
	headerBPCustomer := psdc.HeaderBPCustomer
	headerBPSupplier := psdc.HeaderBPSupplier

	if *headerBPCustomer.TransactionCurrency != *headerBPSupplier.TransactionCurrency {
		return xerrors.Errorf("得意先データと仕入先データのCurrencyが一致しません。")
	}

	return nil
}

func (f *SubFunction) TransactionCurrency(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.TransactionCurrency, error) {

	data := psdc.ConvertToTransactionCurrency()

	return data, nil
}

func GetDateStr() string {
	day := time.Now()
	return day.Format("2006-01-02")
}

func (f *SubFunction) TotalNetAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.TotalNetAmount, error) {
	var totalNetAmount float32 = 0

	netAmount := psdc.NetAmount

	for _, v := range *netAmount {
		if v.NetAmount != nil {
			totalNetAmount += *v.NetAmount
		}
	}

	if sdc.Orders.TotalNetAmount != nil {
		if *sdc.Orders.TotalNetAmount != totalNetAmount {
			return nil, xerrors.Errorf("入力ファイルのTotalNetAmountと計算結果が一致しません。")
		}
	}

	data := psdc.ConvertToTotalNetAmount(totalNetAmount)

	return data, nil
}

func (f *SubFunction) TotalTaxAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.TotalTaxAmount, error) {
	var totalTaxAmount float32 = 0

	taxAmount := psdc.TaxAmount

	for _, v := range *taxAmount {
		if v.TaxAmount != nil {
			totalTaxAmount += *v.TaxAmount
		}
	}

	if sdc.Orders.TotalTaxAmount != nil {
		if *sdc.Orders.TotalTaxAmount != totalTaxAmount {
			return nil, xerrors.Errorf("入力ファイルのTotalTaxAmountと計算結果が一致しません。")
		}
	}

	data := psdc.ConvertToTotalTaxAmount(totalTaxAmount)

	return data, nil
}

func (f *SubFunction) TotalGrossAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.TotalGrossAmount, error) {
	var totalGrossAmount float32 = 0

	grossAmount := psdc.GrossAmount

	for _, v := range *grossAmount {
		if v.GrossAmount != nil {
			totalGrossAmount += *v.GrossAmount
		}
	}

	if sdc.Orders.TotalGrossAmount != nil {
		if *sdc.Orders.TotalGrossAmount != totalGrossAmount {
			return nil, xerrors.Errorf("入力ファイルのTotalGrossAmountと計算結果が一致しません。")
		}
	}

	data := psdc.ConvertToTotalGrossAmount(totalGrossAmount)

	return data, nil
}
