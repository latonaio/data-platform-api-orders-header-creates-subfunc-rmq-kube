package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
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
			`SELECT BusinessPartner, Customer, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_data
		WHERE (BusinessPartner, Customer) = (?, ?);`, buyerSellerDetection.BusinessPartnerID, buyerSellerDetection.Buyer,
		)
		if err != nil {
			return nil, err
		}
	} else if buyerSellerDetection.BuyerOrSeller == "Buyer" {
		rows, err = f.db.Query(
			`SELECT BusinessPartner, Supplier, Incoterms, PaymentTerms, PaymentMethod, BPAccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data
		WHERE (BusinessPartner, Supplier) = (?, ?);`, buyerSellerDetection.BusinessPartnerID, buyerSellerDetection.Seller,
		)
		if err != nil {
			return nil, err
		}
	}

	data, err := psdc.ConvertToHeaderBPCustomerSupplier(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
