package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-headers-creates-subfunc-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// initializer
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
	pm := &requests.CalculateOrderIDQueryGets{
		ServiceLabel:             "",
		FieldNameWithNumberRange: "",
		OrderIDLatestNumber:      nil,
	}

	for rows.Next() {
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.OrderIDLatestNumber,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
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
	pm := &requests.CalculateOrderID{
		OrderIDLatestNumber: nil,
		OrderID:             nil,
	}

	pm.OrderIDLatestNumber = orderIDLatestNumber
	data := pm

	calculateOrderID := CalculateOrderID{
		OrderIDLatestNumber: data.OrderIDLatestNumber,
		OrderID:             data.OrderID,
	}

	return &calculateOrderID, nil
}

func (psdc *SDC) ConvertToHeaderBPCustomerSupplier(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*HeaderBPCustomerSupplier, error) {
	pm := &requests.HeaderBPCustomerSupplier{
		OrderID:                  nil,
		BusinessPartnerID:        nil,
		CustomerOrSupplier:       nil,
		Incoterms:                "",
		PaymentTerms:             "",
		PaymentMethod:            "",
		BPAccountAssignmentGroup: "",
	}

	for rows.Next() {
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.CustomerOrSupplier,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.BPAccountAssignmentGroup)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerBPCustomerSupplier := HeaderBPCustomerSupplier{
		OrderID:                  data.OrderID,
		BusinessPartnerID:        data.BusinessPartnerID,
		CustomerOrSupplier:       data.CustomerOrSupplier,
		Incoterms:                data.Incoterms,
		PaymentTerms:             data.PaymentTerms,
		PaymentMethod:            data.PaymentMethod,
		BPAccountAssignmentGroup: data.BPAccountAssignmentGroup,
	}

	return &headerBPCustomerSupplier, nil
}

// HeaderPartner
func (psdc *SDC) ConvertToHeaderPartnerFunctionKey() (*HeaderPartnerFunctionKey, error) {
	pm := &requests.HeaderPartnerFunctionKey{
		OrderID:            nil,
		BusinessPartnerID:  nil,
		CustomerOrSupplier: nil,
	}

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

	pm := &requests.HeaderPartnerFunction{
		OrderID:           nil,
		BusinessPartnerID: nil,
		PartnerCounter:    nil,
		PartnerFunction:   "",
		BusinessPartner:   nil,
		DefaultPartner:    nil,
	}

	for rows.Next() {
		err := rows.Scan(
			&pm.BusinessPartnerID,
			&pm.PartnerCounter,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.DefaultPartner)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
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

	pm := &requests.HeaderPartnerBPGeneral{
		OrderID:                 nil,
		BusinessPartner:         nil,
		BusinessPartnerFullName: "",
		BusinessPartnerName:     "",
		Country:                 "",
		Language:                "",
		Currency:                "",
		AddressID:               nil,
	}

	for rows.Next() {
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.AddressID,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
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
			Currency:                data.Currency,
			AddressID:               data.AddressID,
		})
	}

	return &headerPartnerBPGeneral, nil
}

// HeaderPartnerPlant
func (psdc *SDC) ConvertToHeaderPartnerPlantKey(length int) (*[]HeaderPartnerPlantKey, error) {
	var headerPartnerPlantKey []HeaderPartnerPlantKey

	pm := &requests.HeaderPartnerPlantKey{
		OrderID:                        nil,
		BusinessPartnerID:              nil,
		CustomerOrSupplier:             nil,
		PartnerCounter:                 nil,
		PartnerFunction:                "",
		PartnerFunctionBusinessPartner: nil,
	}

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

	pm := &requests.HeaderPartnerPlant{
		OrderID:         nil,
		BusinessPartner: nil,
		PartnerFunction: "",
		PlantCounter:    nil,
		Plant:           "",
		DefaultPlant:    nil,
	}

	for rows.Next() {
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.PartnerFunction,
			&pm.PlantCounter,
			&pm.Plant,
			&pm.DefaultPlant,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
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
