package dpfm_api_output_formatter

import api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"

func ConvertToHeader(
	buyerSellerDetection *api_processing_data_formatter.BuyerSellerDetection,
	calculateOrderID *api_processing_data_formatter.CalculateOrderID,
	headerBPCustomerSupplier *api_processing_data_formatter.HeaderBPCustomerSupplier,
) (*Header, error) {
	pm := &Header{}

	pm.OrderID = calculateOrderID.OrderIDLatestNumber
	pm.Buyer = buyerSellerDetection.Buyer
	pm.Seller = buyerSellerDetection.Seller
	pm.Incoterms = headerBPCustomerSupplier.Incoterms
	pm.PaymentTerms = headerBPCustomerSupplier.PaymentTerms
	pm.PaymentMethod = headerBPCustomerSupplier.PaymentMethod
	pm.BPAccountAssignmentGroup = headerBPCustomerSupplier.BPAccountAssignmentGroup

	data := pm

	header := Header{
		OrderID:                         data.OrderID,
		OrderDate:                       data.OrderDate,
		OrderType:                       data.OrderType,
		Buyer:                           data.Buyer,
		Seller:                          data.Seller,
		CreationDate:                    data.CreationDate,
		LastChangeDate:                  data.LastChangeDate,
		ContractType:                    data.ContractType,
		ValidityStartDate:               data.ValidityStartDate,
		ValidityEndDate:                 data.ValidityEndDate,
		InvoiceScheduleStartDate:        data.InvoiceScheduleStartDate,
		InvoiceScheduleEndDate:          data.InvoiceScheduleEndDate,
		TotalNetAmount:                  data.TotalNetAmount,
		TotalTaxAmount:                  data.TotalTaxAmount,
		TotalGrossAmount:                data.TotalGrossAmount,
		OverallDeliveryStatus:           data.OverallDeliveryStatus,
		TotalBlockStatus:                data.TotalBlockStatus,
		OverallOrdReltdBillgStatus:      data.OverallOrdReltdBillgStatus,
		OverallDocReferenceStatus:       data.OverallDocReferenceStatus,
		TransactionCurrency:             data.TransactionCurrency,
		PricingDate:                     data.PricingDate,
		PriceDetnExchangeRate:           data.PriceDetnExchangeRate,
		RequestedDeliveryDate:           data.RequestedDeliveryDate,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderBillingBlockReason:        data.HeaderBillingBlockReason,
		DeliveryBlockReason:             data.DeliveryBlockReason,
		Incoterms:                       data.Incoterms,
		PaymentTerms:                    data.PaymentTerms,
		PaymentMethod:                   data.PaymentMethod,
		ReferenceDocument:               data.ReferenceDocument,
		ReferenceDocumentItem:           data.ReferenceDocumentItem,
		BPAccountAssignmentGroup:        data.BPAccountAssignmentGroup,
		AccountingExchangeRate:          data.AccountingExchangeRate,
		BillingDocumentDate:             data.BillingDocumentDate,
		IsExportImportDelivery:          data.IsExportImportDelivery,
		HeaderText:                      data.HeaderText,
	}

	return &header, nil
}

func ConvertToHeaderPartner(
	headerPartnerFunction *[]api_processing_data_formatter.HeaderPartnerFunction,
	headerPartnerBPGeneral *[]api_processing_data_formatter.HeaderPartnerBPGeneral,
) (*[]HeaderPartner, error) {
	var headerPartner []HeaderPartner
	headerPartnerFunctionMap := make(map[int]api_processing_data_formatter.HeaderPartnerFunction, len(*headerPartnerFunction))
	headerPartnerBPGeneralMap := make(map[int]api_processing_data_formatter.HeaderPartnerBPGeneral, len(*headerPartnerBPGeneral))

	for _, v := range *headerPartnerFunction {
		headerPartnerFunctionMap[*v.BusinessPartner] = v
	}

	for _, v := range *headerPartnerBPGeneral {
		headerPartnerBPGeneralMap[*v.BusinessPartner] = v
	}

	pm := &HeaderPartner{}

	for key, _ := range headerPartnerFunctionMap {
		pm.OrderID = headerPartnerFunctionMap[key].OrderID
		pm.PartnerFunction = headerPartnerFunctionMap[key].PartnerFunction
		pm.BusinessPartner = headerPartnerBPGeneralMap[key].BusinessPartner
		pm.BusinessPartnerFullName = headerPartnerBPGeneralMap[key].BusinessPartnerFullName
		pm.BusinessPartnerName = headerPartnerBPGeneralMap[key].BusinessPartnerName
		pm.Country = headerPartnerBPGeneralMap[key].Country
		pm.Language = headerPartnerBPGeneralMap[key].Language
		pm.Currency = headerPartnerBPGeneralMap[key].Currency
		pm.AddressID = headerPartnerBPGeneralMap[key].AddressID

		data := pm
		headerPartner = append(headerPartner, HeaderPartner{
			OrderID:                 data.OrderID,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}

	return &headerPartner, nil
}

func ConvertToHeaderPartnerPlant(
	psdcHeaderPartnerPlant *[]api_processing_data_formatter.HeaderPartnerPlant,
) (*[]HeaderPartnerPlant, error) {
	var headerPartnerPlant []HeaderPartnerPlant

	pm := &HeaderPartnerPlant{}

	for i, _ := range *psdcHeaderPartnerPlant {
		pm.OrderID = (*psdcHeaderPartnerPlant)[i].OrderID
		pm.PartnerFunction = (*psdcHeaderPartnerPlant)[i].PartnerFunction
		pm.BusinessPartner = (*psdcHeaderPartnerPlant)[i].BusinessPartner
		pm.Plant = (*psdcHeaderPartnerPlant)[i].Plant

		data := pm
		headerPartnerPlant = append(headerPartnerPlant, HeaderPartnerPlant{
			OrderID:         data.OrderID,
			PartnerFunction: data.PartnerFunction,
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}

	return &headerPartnerPlant, nil
}
