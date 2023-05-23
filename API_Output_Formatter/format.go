package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*Header, error) {
	calculateOrderID := psdc.CalculateOrderID
	headerBPCustomerSupplier := psdc.HeaderBPCustomerSupplier
	invoiceDocumentDate := psdc.InvoiceDocumentDate
	// paymentDueDate := psdc.PaymentDueDate
	// netPaymentDays := psdc.NetPaymentDays
	transactionCurrency := psdc.TransactionCurrency
	priceDetnExchangeRate := psdc.PriceDetnExchangeRate
	accountingExchangeRate := psdc.AccountingExchangeRate
	totalNetAmount := psdc.TotalNetAmount
	totalTaxAmount := psdc.TotalTaxAmount
	totalGrossAmount := psdc.TotalGrossAmount

	header := Header{}
	inputHeader := sdc.Orders
	inputData, err := json.Marshal(inputHeader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inputData, &header)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(headerBPCustomerSupplier)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &header)
	if err != nil {
		return nil, err
	}

	header.OrderID = calculateOrderID.OrderID
	header.InvoiceDocumentDate = invoiceDocumentDate.InvoiceDocumentDate
	// header.PaymentDueDate = paymentDueDate.PaymentDueDate
	// header.NetPaymentDays = netPaymentDays.NetPaymentDays
	header.TransactionCurrency = transactionCurrency.TransactionCurrency
	header.PriceDetnExchangeRate = priceDetnExchangeRate.PriceDetnExchangeRate
	header.AccountingExchangeRate = accountingExchangeRate.AccountingExchangeRate
	header.TotalNetAmount = totalNetAmount.TotalNetAmount
	header.TotalTaxAmount = totalTaxAmount.TotalTaxAmount
	header.TotalGrossAmount = totalGrossAmount.TotalGrossAmount

	return &header, nil
}

func ConvertToHeaderPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]HeaderPartner, error) {
	var headerPartners []HeaderPartner
	calculateOrderID := psdc.CalculateOrderID
	headerPartnerFunction := psdc.HeaderPartnerFunction
	headerPartnerBPGeneral := psdc.HeaderPartnerBPGeneral
	headerPartnerFunctionMap := make(map[int]api_processing_data_formatter.HeaderPartnerFunction, len(*headerPartnerFunction))
	headerPartnerBPGeneralMap := make(map[int]api_processing_data_formatter.HeaderPartnerBPGeneral, len(*headerPartnerBPGeneral))

	for _, v := range *headerPartnerFunction {
		headerPartnerFunctionMap[*v.BusinessPartner] = v
	}

	for _, v := range *headerPartnerBPGeneral {
		headerPartnerBPGeneralMap[v.BusinessPartner] = v
	}

	for key := range headerPartnerFunctionMap {
		headerPartners = append(headerPartners, HeaderPartner{
			OrderID:                 calculateOrderID.OrderID,
			PartnerFunction:         *headerPartnerFunctionMap[key].PartnerFunction,
			BusinessPartner:         headerPartnerBPGeneralMap[key].BusinessPartner,
			BusinessPartnerFullName: headerPartnerBPGeneralMap[key].BusinessPartnerFullName,
			BusinessPartnerName:     getStringPtr(headerPartnerBPGeneralMap[key].BusinessPartnerName),
			Country:                 getStringPtr(headerPartnerBPGeneralMap[key].Country),
			Language:                getStringPtr(headerPartnerBPGeneralMap[key].Language),
			AddressID:               headerPartnerBPGeneralMap[key].AddressID,
		})
	}

	return &headerPartners, nil
}

func ConvertToHeaderPartnerPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]HeaderPartnerPlant, error) {
	var headerPartnerPlants []HeaderPartnerPlant
	calculateOrderID := psdc.CalculateOrderID
	headerPartnerPlant := psdc.HeaderPartnerPlant

	for i := range *headerPartnerPlant {
		headerPartnerPlants = append(headerPartnerPlants, HeaderPartnerPlant{
			OrderID:         calculateOrderID.OrderID,
			PartnerFunction: (*headerPartnerPlant)[i].PartnerFunction,
			BusinessPartner: (*headerPartnerPlant)[i].BusinessPartner,
			Plant:           *(*headerPartnerPlant)[i].Plant,
		})
	}

	return &headerPartnerPlants, nil
}

func getStringPtr(s string) *string {
	return &s
}
