package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-headers-creates-subfunc-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// initializer
func (psdc *SDC) ConvertToOrderRegistrationType(sdc *api_input_reader.SDC) (*OrderRegistrationType, error) {
	pm := &requests.OrderRegistrationType{
		ReferenceDocument:     sdc.OrdersInputParameters.ReferenceDocument,
		ReferenceDocumentItem: sdc.OrdersInputParameters.ReferenceDocumentItem,
	}
	data := pm

	orderRegistrationType := OrderRegistrationType{
		ReferenceDocument:     data.ReferenceDocument,
		ReferenceDocumentItem: data.ReferenceDocumentItem,
	}

	return &orderRegistrationType, nil
}

func (psdc *SDC) ConvertToOrderReferenceTypeQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderReferenceTypeQueryGets, error) {
	pm := &requests.OrderReferenceTypeQueryGets{}
	var orderReferenceTypeQueryGets []OrderReferenceTypeQueryGets

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_number_range_number_range_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.NumberRangeFrom,
			&pm.NumberRangeTo)
		if err != nil {
			return nil, err
		}

		data := pm
		orderReferenceTypeQueryGets = append(orderReferenceTypeQueryGets, OrderReferenceTypeQueryGets{
			ServiceLabel:             data.ServiceLabel,
			FieldNameWithNumberRange: data.FieldNameWithNumberRange,
			NumberRangeFrom:          data.NumberRangeFrom,
			NumberRangeTo:            data.NumberRangeTo,
		})
	}

	return &orderReferenceTypeQueryGets, nil
}

func (psdc *SDC) ConvertToOrderReferenceType(
	sdc *api_input_reader.SDC,
	orderReferenceTypeQueryGets *OrderReferenceTypeQueryGets,
) (*OrderReferenceType, error) {
	pm := &requests.OrderReferenceType{}

	pm.ServiceLabel = orderReferenceTypeQueryGets.ServiceLabel

	data := pm
	orderReferenceType := OrderReferenceType{
		ServiceLabel: data.ServiceLabel,
	}

	return &orderReferenceType, nil
}

func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) (*MetaData, error) {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}
	data := pm

	metaData := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &metaData, nil
}

func (psdc *SDC) ConvertToBuyerSellerDetection(sdc *api_input_reader.SDC) (*BuyerSellerDetection, error) {
	pm := &requests.BuyerSellerDetection{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
		Buyer:             sdc.Orders.Buyer,
		Seller:            sdc.Orders.Seller,
	}
	data := pm

	buyerSellerDetection := BuyerSellerDetection{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
		Buyer:             data.Buyer,
		Seller:            data.Seller,
	}

	return &buyerSellerDetection, nil
}

// Header
func (psdc *SDC) ConvertToHeaderBPCustomer(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*HeaderBPCustomer, error) {
	pm := &requests.HeaderBPCustomer{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_customer_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.Customer,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.BPAccountAssignmentGroup,
		)
		if err != nil {
			return nil, err
		}
	}
	data := pm

	headerBPCustomer := HeaderBPCustomer{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		Customer:                 data.Customer,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &headerBPCustomer, nil
}

func (psdc *SDC) ConvertToHeaderBPSupplier(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*HeaderBPSupplier, error) {
	pm := &requests.HeaderBPSupplier{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_supplier_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.Supplier,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.BPAccountAssignmentGroup,
		)
		if err != nil {
			return nil, err
		}
	}
	data := pm

	headerBPSupplier := HeaderBPSupplier{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		Supplier:                 data.Supplier,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &headerBPSupplier, nil
}

func (psdc *SDC) ConvertToHeaderBPCustomerSupplier(sdc *api_input_reader.SDC) (*HeaderBPCustomerSupplier, error) {
	pm := &requests.HeaderBPCustomerSupplier{}

	if psdc.BuyerSellerDetection.BuyerOrSeller == "Seller" {
		pm.BusinessPartnerID = psdc.HeaderBPCustomer.BusinessPartnerID
		pm.CustomerOrSupplier = psdc.HeaderBPCustomer.Customer
		pm.TransactionCurrency = psdc.HeaderBPCustomer.TransactionCurrency
		pm.Incoterms = psdc.HeaderBPCustomer.Incoterms
		pm.PaymentTerms = psdc.HeaderBPCustomer.PaymentTerms
		pm.PaymentMethod = psdc.HeaderBPCustomer.PaymentMethod
		pm.BPAccountAssignmentGroup = psdc.HeaderBPCustomer.BPAccountAssignmentGroup
	} else if psdc.BuyerSellerDetection.BuyerOrSeller == "Buyer" {
		pm.BusinessPartnerID = psdc.HeaderBPSupplier.BusinessPartnerID
		pm.CustomerOrSupplier = psdc.HeaderBPSupplier.Supplier
		pm.TransactionCurrency = psdc.HeaderBPSupplier.TransactionCurrency
		pm.Incoterms = psdc.HeaderBPSupplier.Incoterms
		pm.PaymentTerms = psdc.HeaderBPSupplier.PaymentTerms
		pm.PaymentMethod = psdc.HeaderBPSupplier.PaymentMethod
		pm.BPAccountAssignmentGroup = psdc.HeaderBPSupplier.BPAccountAssignmentGroup
	}

	data := pm
	headerBPCustomerSupplier := HeaderBPCustomerSupplier{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		CustomerOrSupplier:       data.CustomerOrSupplier,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &headerBPCustomerSupplier, nil
}

func (psdc *SDC) ConvertToCalculateOrderIDKey() (*CalculateOrderIDKey, error) {
	pm := &requests.CalculateOrderIDKey{
		ServiceLabel:             "",
		FieldNameWithNumberRange: "OrderID",
	}
	data := pm

	calculateOrderIDKey := CalculateOrderIDKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &calculateOrderIDKey, nil
}

func (psdc *SDC) ConvertToCalculateOrderIDQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*CalculateOrderIDQueryGets, error) {
	pm := &requests.CalculateOrderIDQueryGets{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.OrderIDLatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	data := pm

	calculateOrderIDQueryGets := CalculateOrderIDQueryGets{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
		OrderIDLatestNumber:      data.OrderIDLatestNumber,
	}

	return &calculateOrderIDQueryGets, nil
}

func (psdc *SDC) ConvertToCalculateOrderID(
	orderIDLatestNumber *int,
) (*CalculateOrderID, error) {
	pm := &requests.CalculateOrderID{}

	pm.OrderIDLatestNumber = orderIDLatestNumber
	data := pm

	calculateOrderID := CalculateOrderID{
		OrderIDLatestNumber: data.OrderIDLatestNumber,
		OrderID:             data.OrderID,
	}

	return &calculateOrderID, nil
}

func (psdc *SDC) ConvertToPaymentTermsKey() (*PaymentTermsKey, error) {
	pm := &requests.PaymentTermsKey{}

	data := pm
	paymentTermsKey := PaymentTermsKey{
		PaymentTerms: data.PaymentTerms,
	}

	return &paymentTermsKey, nil
}

func (psdc *SDC) ConvertToPaymentTerms(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]PaymentTerms, error) {
	var paymentTerms []PaymentTerms
	pm := &requests.PaymentTerms{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_payment_terms_payment_terms_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.PaymentTerms,
			&pm.BaseDate,
			&pm.BaseDateCalcAddMonth,
			&pm.BaseDateCalcFixedDate,
			&pm.PaymentDueDateCalcAddMonth,
			&pm.PaymentDueDateCalcFixedDate,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		paymentTerms = append(paymentTerms, PaymentTerms{
			PaymentTerms:                data.PaymentTerms,
			BaseDate:                    data.BaseDate,
			BaseDateCalcAddMonth:        data.BaseDateCalcAddMonth,
			BaseDateCalcFixedDate:       data.BaseDateCalcFixedDate,
			PaymentDueDateCalcAddMonth:  data.PaymentDueDateCalcAddMonth,
			PaymentDueDateCalcFixedDate: data.PaymentDueDateCalcFixedDate,
		})
	}
	return &paymentTerms, nil
}

func (psdc *SDC) ConvertToInvoiceDocumentDate(
	sdc *api_input_reader.SDC,
) (*InvoiceDocumentDate, error) {
	pm := &requests.InvoiceDocumentDate{}

	pm.InvoiceDocumentDate = *sdc.Orders.InvoiceDocumentDate

	data := pm
	invoiceDocumentDate := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &invoiceDocumentDate, nil
}

func (psdc *SDC) ConvertToRequestedDeliveryDate(
	sdc *api_input_reader.SDC,
) (*InvoiceDocumentDate, error) {
	pm := &requests.InvoiceDocumentDate{}

	pm.RequestedDeliveryDate = *sdc.Orders.RequestedDeliveryDate

	data := pm
	invoiceDocumentDate := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &invoiceDocumentDate, nil
}

func (psdc *SDC) ConvertToCaluculateInvoiceDocumentDate(
	sdc *api_input_reader.SDC,
	calculateInvoiceDocumentDate string,
) (*InvoiceDocumentDate, error) {
	pm := &requests.InvoiceDocumentDate{}

	pm.RequestedDeliveryDate = *sdc.Orders.RequestedDeliveryDate
	pm.InvoiceDocumentDate = calculateInvoiceDocumentDate

	data := pm
	invoiceDocumentDate := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &invoiceDocumentDate, nil
}

func (psdc *SDC) ConvertToPaymentDueDate(
	sdc *api_input_reader.SDC,
) (*PaymentDueDate, error) {
	pm := &requests.PaymentDueDate{}

	pm.PaymentDueDate = *sdc.Orders.PaymentDueDate

	data := pm
	paymentDueDate := PaymentDueDate{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
	}

	return &paymentDueDate, nil
}

func (psdc *SDC) ConvertToCaluculatePaymentDueDate(
	sdc *api_input_reader.SDC,
	calculatePaymentDueDate string,
) (*PaymentDueDate, error) {
	pm := &requests.PaymentDueDate{}

	pm.InvoiceDocumentDate = psdc.InvoiceDocumentDate.InvoiceDocumentDate
	pm.PaymentDueDate = calculatePaymentDueDate

	data := pm
	paymentDueDate := PaymentDueDate{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
	}

	return &paymentDueDate, nil
}

func (psdc *SDC) ConvertToNetPaymentDays(
	sdc *api_input_reader.SDC,
) (*NetPaymentDays, error) {
	pm := &requests.NetPaymentDays{}

	pm.NetPaymentDays = sdc.Orders.NetPaymentDays

	data := pm
	netPaymentDays := NetPaymentDays{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
		NetPaymentDays:      data.NetPaymentDays,
	}

	return &netPaymentDays, nil
}

func (psdc *SDC) ConvertToCaluculateNetPaymentDays(
	sdc *api_input_reader.SDC,
	calculateNetPaymentDays int,
) (*NetPaymentDays, error) {
	pm := &requests.NetPaymentDays{}

	pm.InvoiceDocumentDate = psdc.InvoiceDocumentDate.InvoiceDocumentDate
	pm.PaymentDueDate = psdc.PaymentDueDate.PaymentDueDate
	pm.NetPaymentDays = &calculateNetPaymentDays

	data := pm
	netPaymentDays := NetPaymentDays{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
		NetPaymentDays:      data.NetPaymentDays,
	}

	return &netPaymentDays, nil
}

func (psdc *SDC) ConvertToOverallDocReferenceStatus(
	serviceLabel string,
) (*OverallDocReferenceStatus, error) {
	pm := &requests.OverallDocReferenceStatus{}

	pm.OverallDocReferenceStatus = serviceLabel

	data := pm
	overallDocReferenceStatus := OverallDocReferenceStatus{
		OverallDocReferenceStatus: data.OverallDocReferenceStatus,
	}

	return &overallDocReferenceStatus, nil
}

func (psdc *SDC) ConvertToPriceDetnExchangeRate(
	inputPriceDetnExchangeRate *float32,
) (*PriceDetnExchangeRate, error) {
	pm := &requests.PriceDetnExchangeRate{}

	pm.PriceDetnExchangeRate = inputPriceDetnExchangeRate

	data := pm
	priceDetnExchangeRate := PriceDetnExchangeRate{
		PriceDetnExchangeRate: data.PriceDetnExchangeRate,
	}

	return &priceDetnExchangeRate, nil
}

func (psdc *SDC) ConvertToAccountingExchangeRate(
	inputAccountingExchangeRate *float32,
) (*AccountingExchangeRate, error) {
	pm := &requests.AccountingExchangeRate{}

	pm.AccountingExchangeRate = inputAccountingExchangeRate

	data := pm
	accountingExchangeRate := AccountingExchangeRate{
		AccountingExchangeRate: data.AccountingExchangeRate,
	}

	return &accountingExchangeRate, nil
}

func (psdc *SDC) ConvertToTransactionCurrency() (*TransactionCurrency, error) {
	pm := &requests.TransactionCurrency{}

	pm.TransactionCurrency = *psdc.HeaderBPCustomerSupplier.TransactionCurrency

	data := pm
	transactionCurrency := TransactionCurrency{
		TransactionCurrency: data.TransactionCurrency,
	}

	return &transactionCurrency, nil
}

// HeaderPartner
func (psdc *SDC) ConvertToHeaderPartnerFunctionKey() (*HeaderPartnerFunctionKey, error) {
	pm := &requests.HeaderPartnerFunctionKey{}

	data := pm
	headerPartnerFunctionKey := HeaderPartnerFunctionKey{
		OrderID:            data.OrderID,
		BusinessPartnerID:  data.BusinessPartnerID,
		CustomerOrSupplier: data.CustomerOrSupplier,
	}

	return &headerPartnerFunctionKey, nil
}

func (psdc *SDC) ConvertToHeaderPartnerFunction(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderPartnerFunction, error) {
	var headerPartnerFunction []HeaderPartnerFunction

	pm := &requests.HeaderPartnerFunction{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_customer_partner_function_data'または'data_platform_business_partner_supplier_partner_function_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.PartnerCounter,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.DefaultPartner)
		if err != nil {
			return nil, err
		}

		data := pm
		headerPartnerFunction = append(headerPartnerFunction, HeaderPartnerFunction{
			OrderID:           data.OrderID,
			BusinessPartnerID: data.BusinessPartnerID,
			PartnerCounter:    data.PartnerCounter,
			PartnerFunction:   data.PartnerFunction,
			BusinessPartner:   data.BusinessPartner,
			DefaultPartner:    data.DefaultPartner,
		})
	}

	return &headerPartnerFunction, nil
}

func (psdc *SDC) ConvertToHeaderPartnerBPGeneral(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderPartnerBPGeneral, error) {
	var headerPartnerBPGeneral []HeaderPartnerBPGeneral

	pm := &requests.HeaderPartnerBPGeneral{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_general_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Country,
			&pm.Language,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerPartnerBPGeneral = append(headerPartnerBPGeneral, HeaderPartnerBPGeneral{
			OrderID:                 data.OrderID,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Country:                 data.Country,
			Language:                data.Language,
			AddressID:               data.AddressID,
		})
	}

	return &headerPartnerBPGeneral, nil
}

// HeaderPartnerPlant
func (psdc *SDC) ConvertToHeaderPartnerPlantKey(length int) (*[]HeaderPartnerPlantKey, error) {
	var headerPartnerPlantKey []HeaderPartnerPlantKey

	pm := &requests.HeaderPartnerPlantKey{}

	data := pm
	for i := 0; i < length; i++ {
		headerPartnerPlantKey = append(headerPartnerPlantKey, HeaderPartnerPlantKey{
			OrderID:                        data.OrderID,
			BusinessPartnerID:              data.BusinessPartnerID,
			CustomerOrSupplier:             data.CustomerOrSupplier,
			PartnerCounter:                 data.PartnerCounter,
			PartnerFunction:                data.PartnerFunction,
			PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		})
	}

	return &headerPartnerPlantKey, nil
}

func (psdc *SDC) ConvertToHeaderPartnerPlant(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderPartnerPlant, error) {
	var headerPartnerPlant []HeaderPartnerPlant

	pm := &requests.HeaderPartnerPlant{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_supplier_partner_plant_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.PartnerFunction,
			&pm.PlantCounter,
			&pm.Plant,
			&pm.DefaultPlant,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerPartnerPlant = append(headerPartnerPlant, HeaderPartnerPlant{
			OrderID:         data.OrderID,
			BusinessPartner: data.BusinessPartner,
			PartnerFunction: data.PartnerFunction,
			PlantCounter:    data.PlantCounter,
			Plant:           data.Plant,
			DefaultPlant:    data.DefaultPlant,
		})
	}

	return &headerPartnerPlant, nil
}
