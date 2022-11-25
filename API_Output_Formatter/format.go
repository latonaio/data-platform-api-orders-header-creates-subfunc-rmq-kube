package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	buyerSellerDetection *api_processing_data_formatter.BuyerSellerDetection,
	calculateOrderID *api_processing_data_formatter.CalculateOrderID,
	headerBPCustomerSupplier *api_processing_data_formatter.HeaderBPCustomerSupplier,
) (*Header, error) {
	header := &Header{
		OrderID:                         calculateOrderID.OrderIDLatestNumber,
		OrderDate:                       sdc.Orders.OrderDate,
		OrderType:                       sdc.Orders.OrderType,
		Buyer:                           sdc.Orders.Buyer,
		Seller:                          sdc.Orders.Seller,
		CreationDate:                    sdc.Orders.CreationDate,
		LastChangeDate:                  sdc.Orders.LastChangeDate,
		ContractType:                    sdc.Orders.ContractType,
		ValidityStartDate:               sdc.Orders.ValidityStartDate,
		ValidityEndDate:                 sdc.Orders.ValidityEndDate,
		InvoiceScheduleStartDate:        sdc.Orders.InvoiceScheduleStartDate,
		InvoiceScheduleEndDate:          sdc.Orders.InvoiceScheduleEndDate,
		TotalNetAmount:                  sdc.Orders.TotalNetAmount,
		TotalTaxAmount:                  sdc.Orders.TotalTaxAmount,
		TotalGrossAmount:                sdc.Orders.TotalGrossAmount,
		OverallDeliveryStatus:           sdc.Orders.OverallDeliveryStatus,
		TotalBlockStatus:                sdc.Orders.TotalBlockStatus,
		OverallOrdReltdBillgStatus:      sdc.Orders.OverallOrdReltdBillgStatus,
		OverallDocReferenceStatus:       sdc.Orders.OverallDocReferenceStatus,
		TransactionCurrency:             sdc.Orders.TransactionCurrency,
		PricingDate:                     sdc.Orders.PricingDate,
		PriceDetnExchangeRate:           sdc.Orders.PriceDetnExchangeRate,
		RequestedDeliveryDate:           sdc.Orders.RequestedDeliveryDate,
		HeaderCompleteDeliveryIsDefined: sdc.Orders.HeaderCompleteDeliveryIsDefined,
		HeaderBillingBlockReason:        sdc.Orders.HeaderBillingBlockReason,
		DeliveryBlockReason:             sdc.Orders.DeliveryBlockReason,
		Incoterms:                       headerBPCustomerSupplier.Incoterms,
		PaymentTerms:                    headerBPCustomerSupplier.PaymentTerms,
		PaymentMethod:                   headerBPCustomerSupplier.PaymentMethod,
		ReferenceDocument:               sdc.Orders.ReferenceDocument,
		ReferenceDocumentItem:           sdc.Orders.ReferenceDocumentItem,
		BPAccountAssignmentGroup:        headerBPCustomerSupplier.BPAccountAssignmentGroup,
		AccountingExchangeRate:          sdc.Orders.AccountingExchangeRate,
		BillingDocumentDate:             sdc.Orders.BillingDocumentDate,
		IsExportImportDelivery:          sdc.Orders.IsExportImportDelivery,
		HeaderText:                      sdc.Orders.HeaderText,
	}

	return header, nil
}

func ConvertToHeaderPartner(
	calculateOrderID *api_processing_data_formatter.CalculateOrderID,
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

	for key := range headerPartnerFunctionMap {
		headerPartner = append(headerPartner, HeaderPartner{
			OrderID:                 calculateOrderID.OrderIDLatestNumber,
			PartnerFunction:         headerPartnerFunctionMap[key].PartnerFunction,
			BusinessPartner:         headerPartnerBPGeneralMap[key].BusinessPartner,
			BusinessPartnerFullName: headerPartnerBPGeneralMap[key].BusinessPartnerFullName,
			BusinessPartnerName:     headerPartnerBPGeneralMap[key].BusinessPartnerName,
			Country:                 headerPartnerBPGeneralMap[key].Country,
			Language:                headerPartnerBPGeneralMap[key].Language,
			Currency:                headerPartnerBPGeneralMap[key].Currency,
			AddressID:               headerPartnerBPGeneralMap[key].AddressID,
		})
	}

	return &headerPartner, nil
}

func ConvertToHeaderPartnerPlant(
	calculateOrderID *api_processing_data_formatter.CalculateOrderID,
	psdcHeaderPartnerPlant *[]api_processing_data_formatter.HeaderPartnerPlant,
) (*[]HeaderPartnerPlant, error) {
	var headerPartnerPlant []HeaderPartnerPlant

	for i := range *psdcHeaderPartnerPlant {
		headerPartnerPlant = append(headerPartnerPlant, HeaderPartnerPlant{
			OrderID:         calculateOrderID.OrderIDLatestNumber,
			PartnerFunction: (*psdcHeaderPartnerPlant)[i].PartnerFunction,
			BusinessPartner: (*psdcHeaderPartnerPlant)[i].BusinessPartner,
			Plant:           (*psdcHeaderPartnerPlant)[i].Plant,
		})
	}

	return &headerPartnerPlant, nil
}
