package subfunction

import (
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (f *SubFunction) ExtractBusinessPartnerCustomerPartnerPlantArray(
	businessPartner *int,
	buyer *int,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
) (models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice, error) {
	where := make([]qm.QueryMod, 0, len(bPCustomerPartnerFunctionArray))
	for i := range bPCustomerPartnerFunctionArray {
		where = append(where,
			qm.Or(
				fmt.Sprintf("(`BusinessPartner`, `Customer`, `PartnerCounter`, `PartnerFunction`, `PartnerFunctionBusinessPartner`) = (%d, %d, %d,'%s',%d)", *businessPartner, *buyer, bPCustomerPartnerFunctionArray[i].PartnerCounter, bPCustomerPartnerFunctionArray[i].PartnerFunction.String, bPCustomerPartnerFunctionArray[i].PartnerFunctionBusinessPartner.Int),
			),
		)
	}

	res, err := models.DataPlatformBusinessPartnerCustomerPartnerPlantData(
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

func (f *SubFunction) ExtractBusinessPartnerSupplierPartnerPlantArray(
	businessPartner *int,
	seller *int,
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
) (models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice, error) {
	where := make([]qm.QueryMod, 0, len(bPSupplierPartnerFunctionArray))
	for i := range bPSupplierPartnerFunctionArray {
		where = append(where,
			qm.Or(
				fmt.Sprintf("(`BusinessPartner`, `Supplier`, `PartnerCounter`, `PartnerFunction`, `PartnerFunctionBusinessPartner`) = (%d, %d, %d,'%s',%d)", *businessPartner, *seller, bPSupplierPartnerFunctionArray[i].PartnerCounter, bPSupplierPartnerFunctionArray[i].PartnerFunction.String, bPSupplierPartnerFunctionArray[i].PartnerFunctionBusinessPartner.Int),
			),
		)
	}

	res, err := models.DataPlatformBusinessPartnerSupplierPartnerPlantData(
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
