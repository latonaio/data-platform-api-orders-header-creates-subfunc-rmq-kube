package api_processing_data_formatter

import (
	"data-platform-api-orders-headers-creates-subfunc-rmq-kube/database/models"
)

func ConvertToHeaderRelatedDataForBuyer(
	headerRelatedData *HeaderRelatedData,
	bPSupplierRecord *models.DataPlatformBusinessPartnerSupplierDatum,
	nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum,
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
	bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice,
) *HeaderRelatedData {
	headerRelatedData.LatestNumber = &nRLatestNumberRecord.LatestNumber.Int
	headerRelatedData.HeaderPartnerRelatedData = ConvertToHeaderPartnerRelatedDataForBuyer(bPSupplierPartnerFunctionArray, bPSupplierPartnerPlantArray)
	return headerRelatedData
}

func ConvertToHeaderPartnerRelatedDataForBuyer(
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
	bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice,
) []HeaderPartnerRelatedData {
	headerPartnerRelatedData := make([]HeaderPartnerRelatedData, 0, len(bPSupplierPartnerFunctionArray))

	bPSupplierPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice, len(bPSupplierPartnerPlantArray))

	for i, v := range bPSupplierPartnerPlantArray {
		bPSupplierPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPSupplierPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPSupplierPartnerPlantArray[i])
	}

	for _, bPSupplierPartnerFunctionRecord := range bPSupplierPartnerFunctionArray {

		bPSupplierPartnerPlantArrayTargeted := bPSupplierPartnerPlantArrayMap[bPSupplierPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int]

		headerPartnerRelatedData = append(headerPartnerRelatedData, HeaderPartnerRelatedData{
			PartnerFunction: ConvertToSupplierPartnerFunctionForBuyer(bPSupplierPartnerFunctionRecord),
			PartnerPlant:    ConvertToSupplierPartnerPlantForBuyer(bPSupplierPartnerPlantArrayTargeted),
		})
	}
	return headerPartnerRelatedData
}

func ConvertToSupplierPartnerFunctionForBuyer(
	bPSupplierPartnerFunctionRecord *models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatum,
) PartnerFunction {
	SupplierPartnerFunctionData := PartnerFunction{
		BusinessPartner: &bPSupplierPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int,
		PartnerCounter:  &bPSupplierPartnerFunctionRecord.PartnerCounter,
		DefaultPartner:  bPSupplierPartnerFunctionRecord.DefaultPartner.Ptr(),
	}
	return SupplierPartnerFunctionData
}

func ConvertToSupplierPartnerPlantForBuyer(
	bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice,
) []PartnerPlant {
	partnerPlant := make([]PartnerPlant, 0, len(bPSupplierPartnerPlantArray))
	for _, bPSupplierPartnerPlantRecord := range bPSupplierPartnerPlantArray {
		partnerPlant = append(partnerPlant, PartnerPlant{
			PlantCounter: &bPSupplierPartnerPlantRecord.PlantCounter,
			DefaultPlant: &bPSupplierPartnerPlantRecord.DefaultPlant.Bool,
		})
	}
	return partnerPlant
}

func ConvertToHeaderRelatedDataForSeller(
	headerRelatedData *HeaderRelatedData,
	bPCustomerRecord *models.DataPlatformBusinessPartnerCustomerDatum,
	nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) *HeaderRelatedData {
	headerRelatedData.LatestNumber = &nRLatestNumberRecord.LatestNumber.Int
	headerRelatedData.HeaderPartnerRelatedData = ConvertToHeaderPartnerRelatedDataForSeller(bPCustomerPartnerFunctionArray, bPCustomerPartnerPlantArray)
	return headerRelatedData
}

func ConvertToHeaderPartnerRelatedDataForSeller(
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) []HeaderPartnerRelatedData {
	headerPartnerRelatedData := make([]HeaderPartnerRelatedData, 0, len(bPCustomerPartnerFunctionArray))

	bPCustomerPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice, len(bPCustomerPartnerPlantArray))

	for i, v := range bPCustomerPartnerPlantArray {
		bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPCustomerPartnerPlantArray[i])
	}

	for _, bPCustomerPartnerFunctionRecord := range bPCustomerPartnerFunctionArray {

		bPCustomerPartnerPlantArrayTargeted := bPCustomerPartnerPlantArrayMap[bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int]

		headerPartnerRelatedData = append(headerPartnerRelatedData, HeaderPartnerRelatedData{
			PartnerFunction: ConvertToCustomerPartnerFunctionForSeller(bPCustomerPartnerFunctionRecord),
			PartnerPlant:    ConvertToCustomerPartnerPlantForSeller(bPCustomerPartnerPlantArrayTargeted),
		})
	}
	return headerPartnerRelatedData
}

func ConvertToCustomerPartnerFunctionForSeller(
	bPCustomerPartnerFunctionRecord *models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatum,
) PartnerFunction {
	customerPartnerFunctionData := PartnerFunction{
		BusinessPartner: &bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int,
		PartnerCounter:  &bPCustomerPartnerFunctionRecord.PartnerCounter,
		DefaultPartner:  bPCustomerPartnerFunctionRecord.DefaultPartner.Ptr(),
	}
	return customerPartnerFunctionData
}

func ConvertToCustomerPartnerPlantForSeller(
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) []PartnerPlant {
	partnerPlant := make([]PartnerPlant, 0, len(bPCustomerPartnerPlantArray))
	for _, bPCustomerPartnerPlantRecord := range bPCustomerPartnerPlantArray {
		partnerPlant = append(partnerPlant, PartnerPlant{
			PlantCounter: &bPCustomerPartnerPlantRecord.PlantCounter,
			DefaultPlant: &bPCustomerPartnerPlantRecord.DefaultPlant.Bool,
		})
	}
	return partnerPlant
}
