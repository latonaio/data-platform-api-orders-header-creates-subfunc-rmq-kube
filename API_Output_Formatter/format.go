package api_output_formatter

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
)

func ConvertToHeaderForBuyer(
	order *api_input_reader.Order,
	bPSupplierRecord *models.DataPlatformBusinessPartnerSupplierDatum,
	nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum,
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
	bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice,
	bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice,
) *api_input_reader.Order {
	order.OrderID = CalculateOrderId(nRLatestNumberRecord.LatestNumber.Int)
	order.Incoterms = bPSupplierRecord.Incoterms.String
	order.PaymentTerms = bPSupplierRecord.PaymentTerms.String
	order.PaymentMethod = bPSupplierRecord.PaymentMethod.String
	order.BPAccountAssignmentGroup = bPSupplierRecord.BPAccountAssignmentGroup.String
	order.HeaderPartner = ConvertToHeaderPartnerForBuyer(order.HeaderPartner, bPSupplierPartnerFunctionArray, bPGeneralArray, bPSupplierPartnerPlantArray)

	return order
}

func ConvertToHeaderPartnerForBuyer(
	inoutHeaderPartner []api_input_reader.HeaderPartner,
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
	bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice,
	bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice,
) []api_input_reader.HeaderPartner {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(bPSupplierPartnerFunctionArray))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(inoutHeaderPartner))
	bPSupplierPartnerFunctionArrayMap := make(map[int]models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatum, len(bPSupplierPartnerFunctionArray))
	bPGeneralArrayMap := make(map[int]models.DataPlatformBusinessPartnerGeneralDatum, len(bPGeneralArray))
	bPSupplierPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice, len(bPSupplierPartnerPlantArray))

	for i, v := range inoutHeaderPartner {
		inoutHeaderPartnerMap[*v.BusinessPartner] = inoutHeaderPartner[i]
	}

	for i, v := range bPSupplierPartnerFunctionArray {
		bPSupplierPartnerFunctionArrayMap[v.PartnerFunctionBusinessPartner.Int] = *bPSupplierPartnerFunctionArray[i]
	}

	for i, v := range bPGeneralArray {
		bPGeneralArrayMap[v.BusinessPartner] = *bPGeneralArray[i]
	}

	for i, v := range bPSupplierPartnerPlantArray {
		bPSupplierPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPSupplierPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPSupplierPartnerPlantArray[i])
	}

	for businessPartnerID, bPSupplierPartnerFunctionRecord := range bPSupplierPartnerFunctionArrayMap {
		bPGeneralRecord := bPGeneralArrayMap[bPSupplierPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int]

		if _, ok := inoutHeaderPartnerMap[businessPartnerID]; !ok {
			inoutHeaderPartnerMap[businessPartnerID] = api_input_reader.HeaderPartner{}
		}

		newHeaderPartner := inoutHeaderPartnerMap[businessPartnerID]

		newHeaderPartner.PartnerFunction = bPSupplierPartnerFunctionRecord.PartnerFunction.String
		newHeaderPartner.BusinessPartner = bPSupplierPartnerFunctionRecord.PartnerFunctionBusinessPartner.Ptr()
		newHeaderPartner.BusinessPartnerFullName = bPGeneralRecord.BusinessPartnerFullName.String
		newHeaderPartner.BusinessPartnerName = bPGeneralRecord.BusinessPartnerName
		newHeaderPartner.Country = bPGeneralRecord.Country
		newHeaderPartner.Language = bPGeneralRecord.Language
		newHeaderPartner.Currency = bPGeneralRecord.Currency
		newHeaderPartner.AddressID = bPGeneralRecord.AddressID.Ptr()

		bPSupplierPartnerPlantArray, ok := bPSupplierPartnerPlantArrayMap[businessPartnerID]
		if ok {
			for i, v := range newHeaderPartner.HeaderPartnerPlant {
				if v.Plant != "" {
					break
				}
				if i == len(newHeaderPartner.HeaderPartnerPlant)-1 {
					newHeaderPartner.HeaderPartnerPlant = nil
				}
			}
			for _, bPSupplierPartnerPlantRecord := range bPSupplierPartnerPlantArray {
				newHeaderPartner.HeaderPartnerPlant = append(newHeaderPartner.HeaderPartnerPlant, api_input_reader.HeaderPartnerPlant{
					Plant: bPSupplierPartnerPlantRecord.Plant.String,
				})
			}
		}

		headerPartners[businessPartnerID] = newHeaderPartner
	}

	res := make([]api_input_reader.HeaderPartner, 0, len(headerPartners))
	for i := range headerPartners {
		res = append(res, headerPartners[i])
	}

	return res
}

func ConvertToHeaderForSeller(
	order *api_input_reader.Order,
	bPCustomerRecord *models.DataPlatformBusinessPartnerCustomerDatum,
	nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) *api_input_reader.Order {
	order.OrderID = CalculateOrderId(nRLatestNumberRecord.LatestNumber.Int)
	order.Incoterms = bPCustomerRecord.Incoterms.String
	order.PaymentTerms = bPCustomerRecord.PaymentTerms.String
	order.PaymentMethod = bPCustomerRecord.PaymentMethod.String
	order.BPAccountAssignmentGroup = bPCustomerRecord.BPAccountAssignmentGroup.String
	order.HeaderPartner = ConvertToHeaderPartnerForSeller(order.HeaderPartner, bPCustomerPartnerFunctionArray, bPGeneralArray, bPCustomerPartnerPlantArray)

	return order
}

func ConvertToHeaderPartnerForSeller(
	inoutHeaderPartner []api_input_reader.HeaderPartner,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) []api_input_reader.HeaderPartner {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(bPCustomerPartnerFunctionArray))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(inoutHeaderPartner))
	bPCustomerPartnerFunctionArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatum, len(bPCustomerPartnerFunctionArray))
	bPGeneralArrayMap := make(map[int]models.DataPlatformBusinessPartnerGeneralDatum, len(bPGeneralArray))
	bPCustomerPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice, len(bPCustomerPartnerPlantArray))

	for i, v := range inoutHeaderPartner {
		inoutHeaderPartnerMap[*v.BusinessPartner] = inoutHeaderPartner[i]
	}

	for i, v := range bPCustomerPartnerFunctionArray {
		bPCustomerPartnerFunctionArrayMap[v.PartnerFunctionBusinessPartner.Int] = *bPCustomerPartnerFunctionArray[i]
	}

	for i, v := range bPGeneralArray {
		bPGeneralArrayMap[v.BusinessPartner] = *bPGeneralArray[i]
	}

	for i, v := range bPCustomerPartnerPlantArray {
		bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPCustomerPartnerPlantArray[i])
	}

	for businessPartnerID, bPCustomerPartnerFunctionRecord := range bPCustomerPartnerFunctionArrayMap {
		bPGeneralRecord := bPGeneralArrayMap[bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int]

		if _, ok := inoutHeaderPartnerMap[businessPartnerID]; !ok {
			inoutHeaderPartnerMap[businessPartnerID] = api_input_reader.HeaderPartner{}
		}

		newHeaderPartner := inoutHeaderPartnerMap[businessPartnerID]

		newHeaderPartner.PartnerFunction = bPCustomerPartnerFunctionRecord.PartnerFunction.String
		newHeaderPartner.BusinessPartner = bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Ptr()
		newHeaderPartner.BusinessPartnerFullName = bPGeneralRecord.BusinessPartnerFullName.String
		newHeaderPartner.BusinessPartnerName = bPGeneralRecord.BusinessPartnerName
		newHeaderPartner.Country = bPGeneralRecord.Country
		newHeaderPartner.Language = bPGeneralRecord.Language
		newHeaderPartner.Currency = bPGeneralRecord.Currency
		newHeaderPartner.AddressID = bPGeneralRecord.AddressID.Ptr()

		bPCustomerPartnerPlantArray, ok := bPCustomerPartnerPlantArrayMap[businessPartnerID]
		if ok {
			for i, v := range newHeaderPartner.HeaderPartnerPlant {
				if v.Plant != "" {
					break
				}
				if i == len(newHeaderPartner.HeaderPartnerPlant)-1 {
					newHeaderPartner.HeaderPartnerPlant = nil
				}
			}
			for _, bPCustomerPartnerPlantRecord := range bPCustomerPartnerPlantArray {
				newHeaderPartner.HeaderPartnerPlant = append(newHeaderPartner.HeaderPartnerPlant, api_input_reader.HeaderPartnerPlant{
					Plant: bPCustomerPartnerPlantRecord.Plant.String,
				})
			}
		}

		headerPartners[businessPartnerID] = newHeaderPartner
	}

	res := make([]api_input_reader.HeaderPartner, 0, len(headerPartners))
	for i := range headerPartners {
		res = append(res, headerPartners[i])
	}

	return res
}

func CalculateOrderId(latestNumber int) *int {
	orderId := latestNumber + 1
	return &orderId
}
