package subfunction

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (f *SubFunction) ExtractBusinessPartnerCustomerPartnerFunctionArray(
	businessPartner *int,
	buyer *int,
) (models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice, error) {
	res, err := models.DataPlatformBusinessPartnerCustomerPartnerFunctionData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Customer=?", *buyer),
	).All(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("sql: no rows in result set")
	}

	return res, nil
}

func (f *SubFunction) ExtractBusinessPartnerSupplierPartnerFunctionArray(
	businessPartner *int,
	seller *int,
) (models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice, error) {
	res, err := models.DataPlatformBusinessPartnerSupplierPartnerFunctionData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Supplier=?", *seller),
	).All(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("sql: no rows in result set")
	}

	return res, nil
}

func (f *SubFunction) ExtractBusinessPartnerGeneralArray(
	headerPartner []api_input_reader.HeaderPartner,
) (models.DataPlatformBusinessPartnerGeneralDatumSlice, error) {
	var res models.DataPlatformBusinessPartnerGeneralDatumSlice
	where := make([]qm.QueryMod, 0, len(headerPartner))
	for _, v := range headerPartner {
		where = append(where,
			qm.Or("BusinessPartner=?", v.BusinessPartner),
		)
	}

	res, err := models.DataPlatformBusinessPartnerGeneralData(
		where...,
	).All(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("sql: no rows in result set")
	}

	return res, nil
}
