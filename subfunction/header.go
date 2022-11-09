package subfunction

import (
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (f *SubFunction) ExtractBusinessPartnerCustomerRecord(businessPartner *int, buyer *int) (*models.DataPlatformBusinessPartnerCustomerDatum, error) {
	res, err := models.DataPlatformBusinessPartnerCustomerData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Customer=?", *buyer),
	).One(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *SubFunction) ExtractBusinessPartnerSupplierRecord(businessPartner *int, seller *int) (*models.DataPlatformBusinessPartnerSupplierDatum, error) {
	res, err := models.DataPlatformBusinessPartnerSupplierData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Supplier=?", *seller),
	).One(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *SubFunction) ExtractNumberRangeLatestNumberRecord(serviceLabel string, property string) (*models.DataPlatformNumberRangeLatestNumberDatum, error) {
	res, err := models.DataPlatformNumberRangeLatestNumberData(
		qm.And("ServiceLabel=?", serviceLabel),
		qm.And("FieldNameWithNumberRange=?", property),
	).One(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	return res, nil
}
