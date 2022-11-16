package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"strings"
)

func (f *SubFunction) HeaderPartnerPlant(
	buyerSellerDetection *api_processing_data_formatter.BuyerSellerDetection,
	headerPartnerFunction *[]api_processing_data_formatter.HeaderPartnerFunction,
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderPartnerPlant, error) {
	var args []interface{}
	var rows *sql.Rows
	var err error

	dataKey, err := psdc.ConvertToHeaderPartnerPlantKey(len(*headerPartnerFunction))
	if err != nil {
		return nil, err
	}

	for i, v := range *headerPartnerFunction {
		(*dataKey)[i].BusinessPartnerID = buyerSellerDetection.BusinessPartnerID
		if psdc.Header.BuyerOrSeller == "Seller" {
			(*dataKey)[i].CustomerOrSupplier = buyerSellerDetection.Buyer
		} else if psdc.Header.BuyerOrSeller == "Buyer" {
			(*dataKey)[i].CustomerOrSupplier = buyerSellerDetection.Seller
		}
		(*dataKey)[i].PartnerCounter = v.PartnerCounter
		(*dataKey)[i].PartnerFunction = v.PartnerFunction
		(*dataKey)[i].PartnerFunctionBusinessPartner = v.BusinessPartner
	}

	repeat := strings.Repeat("(?,?,?,?,?),", len(*dataKey)-1) + "(?,?,?,?,?)"
	for _, tag := range *dataKey {
		args = append(
			args,
			tag.BusinessPartnerID,
			tag.CustomerOrSupplier,
			tag.PartnerCounter,
			tag.PartnerFunction,
			tag.PartnerFunctionBusinessPartner)
	}

	if psdc.Header.BuyerOrSeller == "Seller" {
		rows, err = f.db.Query(
			`SELECT PartnerFunctionBusinessPartner, PartnerFunction, PlantCounter, Plant, DefaultPlant
				FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_partner_plant_data
				WHERE (BusinessPartner, Customer, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner) IN ( `+repeat+` );`, args...,
		)
		if err != nil {
			return nil, err
		}
	} else if psdc.Header.BuyerOrSeller == "Buyer" {
		rows, err = f.db.Query(
			`SELECT PartnerFunctionBusinessPartner, PartnerFunction, PlantCounter, Plant, DefaultPlant
				FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_partner_plant_data
				WHERE (BusinessPartner, Supplier, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner) IN ( `+repeat+` );`, args...,
		)
		if err != nil {
			return nil, err
		}
	}

	data, err := psdc.ConvertToHeaderPartnerPlant(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
