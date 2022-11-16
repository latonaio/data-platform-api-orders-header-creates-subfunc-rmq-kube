package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"
	"fmt"
	"os"
)

func (f *SubFunction) SetValue(
	sdc *api_input_reader.SDC,
	buyerSellerDetection *api_processing_data_formatter.BuyerSellerDetection,
	headerBPCustomerSupplier *api_processing_data_formatter.HeaderBPCustomerSupplier,
	calculateOrderID *api_processing_data_formatter.CalculateOrderID,
	headerPartnerFunction *[]api_processing_data_formatter.HeaderPartnerFunction,
	headerPartnerBPGeneral *[]api_processing_data_formatter.HeaderPartnerBPGeneral,
	headerPartnerPlant *[]api_processing_data_formatter.HeaderPartnerPlant,
) (*api_input_reader.SDC, error) {
	var outHeader *dpfm_api_output_formatter.Header
	var outHeaderPartner *[]dpfm_api_output_formatter.HeaderPartner
	var outHeaderPartnerPlant *[]dpfm_api_output_formatter.HeaderPartnerPlant
	var err error

	outHeader, err = dpfm_api_output_formatter.ConvertToHeader(buyerSellerDetection, calculateOrderID, headerBPCustomerSupplier)
	if err != nil {
		return nil, err
	}
	outHeaderPartner, err = dpfm_api_output_formatter.ConvertToHeaderPartner(headerPartnerFunction, headerPartnerBPGeneral)
	if err != nil {
		return nil, err
	}
	outHeaderPartnerPlant, err = dpfm_api_output_formatter.ConvertToHeaderPartnerPlant(headerPartnerPlant)
	if err != nil {
		return nil, err
	}

	raw, err := json.Marshal(outHeader)
	if err != nil {
		fmt.Printf("data marshal error :%#v", err.Error())
	}
	err = json.Unmarshal(raw, &sdc.Orders)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	raw, err = json.Marshal(outHeaderPartner)
	if err != nil {
		fmt.Printf("data marshal error :%#v", err.Error())
	}
	err = json.Unmarshal(raw, &sdc.Orders.HeaderPartner)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	for i, v := range sdc.Orders.HeaderPartner {
		bp := *v.BusinessPartner
		pf := v.PartnerFunction
		sdc.Orders.HeaderPartner[i].HeaderPartnerPlant = make([]api_input_reader.HeaderPartnerPlant, 0, 1)

		for _, v := range *outHeaderPartnerPlant {
			if *v.BusinessPartner == bp && v.PartnerFunction == pf {
				sdc.Orders.HeaderPartner[i].HeaderPartnerPlant = append(sdc.Orders.HeaderPartner[i].HeaderPartnerPlant, api_input_reader.HeaderPartnerPlant{Plant: v.Plant})
			}
		}
	}

	return sdc, nil
}
