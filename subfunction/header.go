package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"sort"
	"time"

	"golang.org/x/xerrors"
)

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
		psdc.HeaderBPCustomer, err = psdc.ConvertToHeaderBPCustomer(sdc, rows)
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
		psdc.HeaderBPSupplier, err = psdc.ConvertToHeaderBPSupplier(sdc, rows)
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
		psdc.HeaderBPSupplier, err = psdc.ConvertToHeaderBPSupplier(sdc, rows)
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
		psdc.HeaderBPCustomer, err = psdc.ConvertToHeaderBPCustomer(sdc, rows)
		if err != nil {
			return nil, err
		}
	}

	data, err := psdc.ConvertToHeaderBPCustomerSupplier(sdc)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CalculateOrderID(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateOrderID, error) {
	metaData := psdc.MetaData
	dataKey, err := psdc.ConvertToCalculateOrderIDKey()
	if err != nil {
		return nil, err
	}

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateOrderIDQueryGets(sdc, rows)
	if err != nil {
		return nil, err
	}

	calculateOrderID := CalculateOrderID(*dataQueryGets.OrderIDLatestNumber)

	data, err := psdc.ConvertToCalculateOrderID(calculateOrderID)
	if err != nil {
		return nil, err
	}

	return data, err
}

func CalculateOrderID(latestNumber int) *int {
	res := latestNumber + 1
	return &res
}

func (f *SubFunction) InvoiceDocumentDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.InvoiceDocumentDate, error) {

	dataKey, err := psdc.ConvertToPaymentTermsKey()
	if err != nil {
		return nil, err
	}

	dataKey.PaymentTerms = psdc.HeaderBPCustomerSupplier.PaymentTerms

	rows, err := f.db.Query(
		`SELECT PaymentTerms, BaseDate, BaseDateCalcAddMonth, BaseDateCalcFixedDate, PaymentDueDateCalcAddMonth, PaymentDueDateCalcFixedDate
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data
			WHERE PaymentTerms = ?;`, dataKey.PaymentTerms,
	)
	if err != nil {
		return nil, err
	}

	psdc.PaymentTerms, err = psdc.ConvertToPaymentTerms(sdc, rows)
	if err != nil {
		return nil, err
	}

	if sdc.Orders.InvoiceDocumentDate != nil {
		if *sdc.Orders.InvoiceDocumentDate != "" {
			data, err := psdc.ConvertToInvoiceDocumentDate(sdc)
			if err != nil {
				return nil, err
			}
			return data, nil
		}
	}

	requestedDeliveryDate, err := psdc.ConvertToRequestedDeliveryDate(sdc)
	if err != nil {
		return nil, err
	}

	calculateInvoiceDocumentDate, err := CalculateInvoiceDocumentDate(psdc, requestedDeliveryDate.RequestedDeliveryDate, psdc.PaymentTerms)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToCaluculateInvoiceDocumentDate(sdc, calculateInvoiceDocumentDate)
	if err != nil {
		return nil, err
	}

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
			data, err := psdc.ConvertToPaymentDueDate(sdc)
			if err != nil {
				return nil, err
			}
			return data, nil
		}
	}

	calculatePaymentDueDate, err := CalculatePaymentDueDate(psdc.InvoiceDocumentDate.InvoiceDocumentDate, psdc.PaymentTerms)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToCaluculatePaymentDueDate(sdc, calculatePaymentDueDate)
	if err != nil {
		return nil, err
	}

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
		if day == *v.BaseDateCalcFixedDate {
			t = time.Date(t.Year(), t.Month()+time.Month(*v.PaymentDueDateCalcAddMonth)+1, 0, 0, 0, 0, 0, time.UTC)
			if *v.PaymentDueDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.PaymentDueDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
			break
		}
		if i == len(*paymentTerms)-1 {
			return "", xerrors.Errorf("入力ファイルの'InvoiceDocumentDate'の値が不適切です。")
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
		data, err := psdc.ConvertToNetPaymentDays(sdc)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	calculateNetPaymentDays, err := CalculateNetPaymentDays(psdc.InvoiceDocumentDate.InvoiceDocumentDate, psdc.PaymentDueDate.PaymentDueDate)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToCaluculateNetPaymentDays(sdc, calculateNetPaymentDays)
	if err != nil {
		return nil, err
	}

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
	res := psdc.OrderReferenceType.ServiceLabel
	data := &api_processing_data_formatter.OverallDocReferenceStatus{}
	var err error

	if res != nil {
		if *res == "QUOTATIONS" {
			data, err = psdc.ConvertToOverallDocReferenceStatus("QT")
			if err != nil {
				return nil, err
			}
		} else if *res == "INQUIRIES" {
			data, err = psdc.ConvertToOverallDocReferenceStatus("IN")
			if err != nil {
				return nil, err
			}
		} else if *res == "PURCHASE_REQUISITION" {
			data, err = psdc.ConvertToOverallDocReferenceStatus("PR")
			if err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}

func (f *SubFunction) PriceDetnExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.PriceDetnExchangeRate, error) {
	inputPriceDetnExchangeRate := sdc.Orders.PriceDetnExchangeRate
	data := &api_processing_data_formatter.PriceDetnExchangeRate{}
	var err error

	if inputPriceDetnExchangeRate != nil {
		data, err = psdc.ConvertToPriceDetnExchangeRate(inputPriceDetnExchangeRate)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (f *SubFunction) AccountingExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.AccountingExchangeRate, error) {
	inputAccountingExchangeRate := sdc.Orders.AccountingExchangeRate
	data := &api_processing_data_formatter.AccountingExchangeRate{}
	var err error

	if inputAccountingExchangeRate != nil {
		data, err = psdc.ConvertToAccountingExchangeRate(inputAccountingExchangeRate)
		if err != nil {
			return nil, err
		}
	}

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

	data, err := psdc.ConvertToTransactionCurrency()
	if err != nil {
		return nil, err
	}

	return data, nil
}
