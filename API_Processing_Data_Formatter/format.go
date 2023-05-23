package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-headers-creates-subfunc-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}

	data := pm
	res := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderRegistrationType(sdc *api_input_reader.SDC) *OrderRegistrationType {
	pm := &requests.OrderRegistrationType{
		ReferenceDocument:     sdc.OrdersInputParameters.ReferenceDocument,
		ReferenceDocumentItem: sdc.OrdersInputParameters.ReferenceDocumentItem,
		RegistrationType:      "直接登録",
	}

	if pm.ReferenceDocument != nil {
		if pm.ReferenceDocumentItem != nil {
			pm.RegistrationType = "参照登録"
		}
	}

	data := pm
	res := OrderRegistrationType{
		ReferenceDocument:     data.ReferenceDocument,
		ReferenceDocumentItem: data.ReferenceDocumentItem,
		RegistrationType:      data.RegistrationType,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderReferenceTypeQueryGets(rows *sql.Rows) (*[]OrderReferenceTypeQueryGets, error) {
	var res []OrderReferenceTypeQueryGets

	for i := 0; true; i++ {
		pm := &requests.OrderReferenceTypeQueryGets{}
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
		res = append(res, OrderReferenceTypeQueryGets{
			ServiceLabel:             data.ServiceLabel,
			FieldNameWithNumberRange: data.FieldNameWithNumberRange,
			NumberRangeFrom:          data.NumberRangeFrom,
			NumberRangeTo:            data.NumberRangeTo,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToOrderReferenceType(orderReferenceTypeQueryGets *OrderReferenceTypeQueryGets) *OrderReferenceType {
	pm := &requests.OrderReferenceType{}

	pm.ServiceLabel = orderReferenceTypeQueryGets.ServiceLabel

	data := pm
	res := OrderReferenceType{
		ServiceLabel: data.ServiceLabel,
	}

	return &res
}

// Header
func (psdc *SDC) ConvertToBuyerSellerDetection(sdc *api_input_reader.SDC, buyerOrSeller string) *BuyerSellerDetection {
	pm := &requests.BuyerSellerDetection{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
		Buyer:             sdc.Orders.Buyer,
		Seller:            sdc.Orders.Seller,
	}

	pm.BuyerOrSeller = buyerOrSeller

	data := pm
	res := BuyerSellerDetection{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
		Buyer:             data.Buyer,
		Seller:            data.Seller,
		BuyerOrSeller:     data.BuyerOrSeller,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderBPCustomer(rows *sql.Rows) (*HeaderBPCustomer, error) {
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

	res := HeaderBPCustomer{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		Customer:                 data.Customer,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToHeaderBPSupplier(rows *sql.Rows) (*HeaderBPSupplier, error) {
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
	res := HeaderBPSupplier{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		Supplier:                 data.Supplier,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToHeaderBPCustomerSupplier() *HeaderBPCustomerSupplier {
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
	res := HeaderBPCustomerSupplier{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		CustomerOrSupplier:       data.CustomerOrSupplier,
		TransactionCurrency:      data.TransactionCurrency,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &res
}

func (psdc *SDC) ConvertToCalculateOrderIDKey() *CalculateOrderIDKey {
	pm := &requests.CalculateOrderIDKey{
		FieldNameWithNumberRange: "OrderID",
	}

	data := pm
	res := CalculateOrderIDKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &res
}

func (psdc *SDC) ConvertToCalculateOrderIDQueryGets(rows *sql.Rows) (*CalculateOrderIDQueryGets, error) {
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
	res := CalculateOrderIDQueryGets{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
		OrderIDLatestNumber:      data.OrderIDLatestNumber,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCalculateOrderID(orderIDLatestNumber *int, orderID int) *CalculateOrderID {
	pm := &requests.CalculateOrderID{}

	pm.OrderIDLatestNumber = orderIDLatestNumber
	pm.OrderID = orderID

	data := pm
	res := CalculateOrderID{
		OrderIDLatestNumber: data.OrderIDLatestNumber,
		OrderID:             data.OrderID,
	}

	return &res
}

func (psdc *SDC) ConvertToPaymentTermsKey() *PaymentTermsKey {
	pm := &requests.PaymentTermsKey{}

	data := pm
	res := PaymentTermsKey{
		PaymentTerms: data.PaymentTerms,
	}

	return &res
}

func (psdc *SDC) ConvertToPaymentTerms(rows *sql.Rows) (*[]PaymentTerms, error) {
	var res []PaymentTerms

	for i := 0; true; i++ {
		pm := &requests.PaymentTerms{}
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
		res = append(res, PaymentTerms{
			PaymentTerms:                data.PaymentTerms,
			BaseDate:                    data.BaseDate,
			BaseDateCalcAddMonth:        data.BaseDateCalcAddMonth,
			BaseDateCalcFixedDate:       data.BaseDateCalcFixedDate,
			PaymentDueDateCalcAddMonth:  data.PaymentDueDateCalcAddMonth,
			PaymentDueDateCalcFixedDate: data.PaymentDueDateCalcFixedDate,
		})
	}
	return &res, nil
}

func (psdc *SDC) ConvertToInvoiceDocumentDate(sdc *api_input_reader.SDC) *InvoiceDocumentDate {
	pm := &requests.InvoiceDocumentDate{
		InvoiceDocumentDate: *sdc.Orders.InvoiceDocumentDate,
	}

	data := pm
	res := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res
}

func (psdc *SDC) ConvertToRequestedDeliveryDate(sdc *api_input_reader.SDC) *InvoiceDocumentDate {
	pm := &requests.InvoiceDocumentDate{
		RequestedDeliveryDate: *sdc.Orders.RequestedDeliveryDate,
	}

	data := pm
	res := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res
}

func (psdc *SDC) ConvertToCaluculateInvoiceDocumentDate(sdc *api_input_reader.SDC, calculateInvoiceDocumentDate string) *InvoiceDocumentDate {
	pm := &requests.InvoiceDocumentDate{
		RequestedDeliveryDate: *sdc.Orders.RequestedDeliveryDate,
	}

	pm.InvoiceDocumentDate = calculateInvoiceDocumentDate

	data := pm
	res := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPaymentDueDate(sdc *api_input_reader.SDC) *PaymentDueDate {
	pm := &requests.PaymentDueDate{
		PaymentDueDate: *sdc.Orders.PaymentDueDate,
	}

	data := pm
	res := PaymentDueDate{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
	}

	return &res
}

func (psdc *SDC) ConvertToCaluculatePaymentDueDate(calculatePaymentDueDate string) *PaymentDueDate {
	pm := &requests.PaymentDueDate{}

	pm.InvoiceDocumentDate = psdc.InvoiceDocumentDate.InvoiceDocumentDate
	pm.PaymentDueDate = calculatePaymentDueDate

	data := pm
	res := PaymentDueDate{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
	}

	return &res
}

func (psdc *SDC) ConvertToNetPaymentDays(sdc *api_input_reader.SDC) *NetPaymentDays {
	pm := &requests.NetPaymentDays{
		NetPaymentDays: sdc.Orders.NetPaymentDays,
	}

	data := pm
	res := NetPaymentDays{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
		NetPaymentDays:      data.NetPaymentDays,
	}

	return &res
}

func (psdc *SDC) ConvertToCaluculateNetPaymentDays(calculateNetPaymentDays int) *NetPaymentDays {
	pm := &requests.NetPaymentDays{}

	pm.InvoiceDocumentDate = psdc.InvoiceDocumentDate.InvoiceDocumentDate
	pm.PaymentDueDate = psdc.PaymentDueDate.PaymentDueDate
	pm.NetPaymentDays = &calculateNetPaymentDays

	data := pm
	res := NetPaymentDays{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
		NetPaymentDays:      data.NetPaymentDays,
	}

	return &res
}

func (psdc *SDC) ConvertToOverallDocReferenceStatus(overallDocReferenceStatus string) *OverallDocReferenceStatus {
	pm := &requests.OverallDocReferenceStatus{}

	pm.OverallDocReferenceStatus = overallDocReferenceStatus

	data := pm
	res := OverallDocReferenceStatus{
		OverallDocReferenceStatus: data.OverallDocReferenceStatus,
	}

	return &res
}

func (psdc *SDC) ConvertToPricingDate(inputPricingDate string) *PricingDate {
	pm := &requests.PricingDate{}

	pm.PricingDate = inputPricingDate

	data := pm
	res := PricingDate{
		PricingDate: data.PricingDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPriceDetnExchangeRate(sdc *api_input_reader.SDC) *PriceDetnExchangeRate {
	pm := &requests.PriceDetnExchangeRate{
		PriceDetnExchangeRate: sdc.Orders.PriceDetnExchangeRate,
	}

	data := pm
	res := PriceDetnExchangeRate{
		PriceDetnExchangeRate: data.PriceDetnExchangeRate,
	}

	return &res
}

func (psdc *SDC) ConvertToAccountingExchangeRate(sdc *api_input_reader.SDC) *AccountingExchangeRate {
	pm := &requests.AccountingExchangeRate{
		AccountingExchangeRate: sdc.Orders.AccountingExchangeRate,
	}

	data := pm
	res := AccountingExchangeRate{
		AccountingExchangeRate: data.AccountingExchangeRate,
	}

	return &res
}

func (psdc *SDC) ConvertToTransactionCurrency() *TransactionCurrency {
	pm := &requests.TransactionCurrency{}

	pm.TransactionCurrency = *psdc.HeaderBPCustomerSupplier.TransactionCurrency

	data := pm
	res := TransactionCurrency{
		TransactionCurrency: data.TransactionCurrency,
	}

	return &res
}

func (psdc *SDC) ConvertToTotalNetAmount(totalNetAmount float32) *TotalNetAmount {
	pm := &requests.TotalNetAmount{}

	pm.TotalNetAmount = totalNetAmount

	data := pm
	res := TotalNetAmount{
		TotalNetAmount: data.TotalNetAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToTotalTaxAmount(totalTaxAmount float32) *TotalTaxAmount {
	pm := &requests.TotalTaxAmount{}

	pm.TotalTaxAmount = totalTaxAmount

	data := pm
	res := TotalTaxAmount{
		TotalTaxAmount: data.TotalTaxAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToTotalGrossAmount(totalGrossAmount float32) *TotalGrossAmount {
	pm := &requests.TotalGrossAmount{}

	pm.TotalGrossAmount = totalGrossAmount

	data := pm
	res := TotalGrossAmount{
		TotalGrossAmount: data.TotalGrossAmount,
	}

	return &res
}

// HeaderPartner
func (psdc *SDC) ConvertToHeaderPartnerFunctionKey() *HeaderPartnerFunctionKey {
	pm := &requests.HeaderPartnerFunctionKey{}

	data := pm
	res := HeaderPartnerFunctionKey{
		BusinessPartnerID:  data.BusinessPartnerID,
		CustomerOrSupplier: data.CustomerOrSupplier,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderPartnerFunction(rows *sql.Rows) (*[]HeaderPartnerFunction, error) {
	var res []HeaderPartnerFunction

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerFunction{}
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
		res = append(res, HeaderPartnerFunction{
			BusinessPartnerID: data.BusinessPartnerID,
			PartnerCounter:    data.PartnerCounter,
			PartnerFunction:   data.PartnerFunction,
			BusinessPartner:   data.BusinessPartner,
			DefaultPartner:    data.DefaultPartner,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToHeaderPartnerBPGeneral(rows *sql.Rows) (*[]HeaderPartnerBPGeneral, error) {
	var res []HeaderPartnerBPGeneral

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerBPGeneral{}
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
		res = append(res, HeaderPartnerBPGeneral{
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Country:                 data.Country,
			Language:                data.Language,
			AddressID:               data.AddressID,
		})
	}

	return &res, nil
}

// HeaderPartnerPlant
func (psdc *SDC) ConvertToHeaderPartnerPlantKey(length int) *[]HeaderPartnerPlantKey {
	var res []HeaderPartnerPlantKey

	for i := 0; i < length; i++ {
		pm := &requests.HeaderPartnerPlantKey{}

		data := pm
		res = append(res, HeaderPartnerPlantKey{
			BusinessPartnerID:              data.BusinessPartnerID,
			CustomerOrSupplier:             data.CustomerOrSupplier,
			PartnerCounter:                 data.PartnerCounter,
			PartnerFunction:                data.PartnerFunction,
			PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderPartnerPlant(rows *sql.Rows) (*[]HeaderPartnerPlant, error) {
	var res []HeaderPartnerPlant

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerPlant{}
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
			&pm.DefaultStockConfirmationPlant,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, HeaderPartnerPlant{
			BusinessPartner:               data.BusinessPartner,
			PartnerFunction:               data.PartnerFunction,
			PlantCounter:                  data.PlantCounter,
			Plant:                         data.Plant,
			DefaultPlant:                  data.DefaultPlant,
			DefaultStockConfirmationPlant: data.DefaultStockConfirmationPlant,
		})
	}

	return &res, nil
}

// Item
func (psdc *SDC) ConvertToItemBPTaxClassificationKey() *ItemBPTaxClassificationKey {
	pm := &requests.ItemBPTaxClassificationKey{}

	data := pm
	res := ItemBPTaxClassificationKey{
		BusinessPartnerID:  data.BusinessPartnerID,
		CustomerOrSupplier: data.CustomerOrSupplier,
		DepartureCountry:   data.DepartureCountry,
	}

	return &res
}

func (psdc *SDC) ConvertToItemBPTaxClassification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*ItemBPTaxClassification, error) {
	pm := &requests.ItemBPTaxClassification{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_customer_tax_data'または'data_platform_business_partner_supplier_tax_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.CustomerOrSupplier,
			&pm.DepartureCountry,
			&pm.BPTaxClassification,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	res := ItemBPTaxClassification{
		BusinessPartnerID:   data.BusinessPartnerID,
		CustomerOrSupplier:  data.CustomerOrSupplier,
		DepartureCountry:    data.DepartureCountry,
		BPTaxClassification: data.BPTaxClassification,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToItemProductTaxClassificationKey(length int) *[]ItemProductTaxClassificationKey {
	var res []ItemProductTaxClassificationKey

	for i := 0; i < length; i++ {
		pm := &requests.ItemProductTaxClassificationKey{
			TaxCategory: "MWST",
		}

		data := pm
		res = append(res, ItemProductTaxClassificationKey{
			Product:           data.Product,
			BusinessPartnerID: data.BusinessPartnerID,
			Country:           data.Country,
			TaxCategory:       data.TaxCategory,
		})
	}
	return &res
}

func (psdc *SDC) ConvertToItemProductTaxClassification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]ItemProductTaxClassification, error) {
	var itemProductTaxClassification []ItemProductTaxClassification

	for i := 0; true; i++ {
		pm := &requests.ItemProductTaxClassification{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_tax_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartnerID,
			&pm.Country,
			&pm.TaxCategory,
			&pm.ProductTaxClassification,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemProductTaxClassification = append(itemProductTaxClassification, ItemProductTaxClassification{
			Product:                  data.Product,
			BusinessPartnerID:        data.BusinessPartnerID,
			Country:                  data.Country,
			TaxCategory:              data.TaxCategory,
			ProductTaxClassification: data.ProductTaxClassification,
		})
	}

	return &itemProductTaxClassification, nil
}

func (psdc *SDC) ConvertToTaxCode(bpTaxClassification, productTaxClassification *string, product, orderType, taxCode string) *TaxCode {
	pm := requests.TaxCode{}

	pm.Product = product
	pm.BPTaxClassification = bpTaxClassification
	pm.ProductTaxClassification = productTaxClassification
	pm.OrderType = orderType
	pm.TaxCode = taxCode

	data := pm
	res := TaxCode{
		Product:                  data.Product,
		BPTaxClassification:      data.BPTaxClassification,
		ProductTaxClassification: data.ProductTaxClassification,
		OrderType:                data.OrderType,
		TaxCode:                  data.TaxCode,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxRateKey() *TaxRateKey {
	pm := &requests.TaxRateKey{
		Country: "JP",
	}

	data := pm
	res := TaxRateKey{
		Country:           data.Country,
		TaxCode:           data.TaxCode,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxRate(rows *sql.Rows) (*[]TaxRate, error) {
	var res []TaxRate

	for i := 0; true; i++ {
		pm := &requests.TaxRate{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_price_master_price_master_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Country,
			&pm.TaxCode,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.TaxRate,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		res = append(res, TaxRate{
			Country:           data.Country,
			TaxCode:           data.TaxCode,
			ValidityEndDate:   data.ValidityEndDate,
			ValidityStartDate: data.ValidityStartDate,
			TaxRate:           data.TaxRate,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToNetAmount(conditionAmount *[]ConditionAmount) *[]NetAmount {
	var res []NetAmount

	for _, v := range *conditionAmount {
		pm := &requests.NetAmount{}

		pm.Product = v.Product
		pm.NetAmount = v.ConditionAmount

		data := pm
		res = append(res, NetAmount{
			Product:   data.Product,
			NetAmount: data.NetAmount,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToTaxAmount(product, taxCode string, taxRate, netAmount, taxAmount *float32) *TaxAmount {
	pm := &requests.TaxAmount{}

	pm.Product = product
	pm.TaxCode = taxCode
	pm.TaxRate = taxRate
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount

	data := pm
	res := TaxAmount{
		Product:   data.Product,
		TaxCode:   data.TaxCode,
		TaxRate:   data.TaxRate,
		NetAmount: data.NetAmount,
		TaxAmount: data.TaxAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToGrossAmount(product string, netAmount, taxAmount, grossAmount *float32) *GrossAmount {
	pm := &requests.GrossAmount{}

	pm.Product = product
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount
	pm.GrossAmount = grossAmount

	data := pm
	res := GrossAmount{
		Product:     data.Product,
		NetAmount:   data.NetAmount,
		TaxAmount:   data.TaxAmount,
		GrossAmount: data.GrossAmount,
	}

	return &res
}

// ItemPricingElement
func (psdc *SDC) ConvertToPriceMasterKey(sdc *api_input_reader.SDC) *PriceMasterKey {
	pm := &requests.PriceMasterKey{
		BusinessPartnerID: sdc.BusinessPartnerID,
	}

	data := pm
	res := PriceMasterKey{
		BusinessPartnerID:          data.BusinessPartnerID,
		Product:                    data.Product,
		CustomerOrSupplier:         data.CustomerOrSupplier,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPriceMaster(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]PriceMaster, error) {
	var res []PriceMaster

	for i := 0; true; i++ {
		pm := &requests.PriceMaster{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_price_master_price_master_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.Product,
			&pm.CustomerOrSupplier,
			&pm.ConditionValidityEndDate,
			&pm.ConditionValidityStartDate,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.ConditionRateValue,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		res = append(res, PriceMaster{
			BusinessPartnerID:          data.BusinessPartnerID,
			Product:                    data.Product,
			CustomerOrSupplier:         data.CustomerOrSupplier,
			ConditionValidityEndDate:   data.ConditionValidityEndDate,
			ConditionValidityStartDate: data.ConditionValidityStartDate,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			ConditionRateValue:         data.ConditionRateValue,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToConditionAmount(product string, conditionQuantity *float32, conditionAmount *float32) *ConditionAmount {
	pm := &requests.ConditionAmount{
		ConditionIsManuallyChanged: GetBoolPtr(false),
	}

	pm.Product = product
	pm.ConditionQuantity = conditionQuantity
	pm.ConditionAmount = conditionAmount

	data := pm
	res := ConditionAmount{
		Product:                    data.Product,
		ConditionQuantity:          data.ConditionQuantity,
		ConditionAmount:            data.ConditionAmount,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}

	return &res
}

func GetBoolPtr(b bool) *bool {
	return &b
}
