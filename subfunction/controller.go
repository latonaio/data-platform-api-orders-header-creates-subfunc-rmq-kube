package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	api_output_formatter "data-platform-api-orders-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
	"fmt"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *SubFunction) Controller(sdc *api_input_reader.SDC) error {
	var err error

	businessPartner := sdc.BusinessPartner
	serviceLabel := sdc.ServiceLabel
	property := "OrderID"
	buyer := sdc.Orders.Buyer
	seller := sdc.Orders.Seller
	processingData := &api_processing_data_formatter.HeaderRelatedData{}

	// 1-0. 入力ファイルのbusiness_partnerがBuyerであるかSellerであるかの判断
	if *businessPartner == *buyer && *businessPartner != *seller {
		processingData.BuyerOrSeller = "Buyer"
		f.l.Info(processingData.BuyerOrSeller)
		err = f.CreateSdcForBuyer(sdc, businessPartner, serviceLabel, property, seller, processingData)
		return err
	} else if *businessPartner != *buyer && *businessPartner == *seller {
		processingData.BuyerOrSeller = "Seller"
		f.l.Info(processingData.BuyerOrSeller)
		err = f.CreateSdcForSeller(sdc, businessPartner, serviceLabel, property, buyer, processingData)
		return err
	} else {
		return fmt.Errorf("business_partnerがBuyerまたはSellerと一致しません")
	}
}

func (f *SubFunction) CreateSdcForBuyer(
	sdc *api_input_reader.SDC,
	businessPartner *int,
	serviceLabel string,
	property string,
	seller *int,
	processingData *api_processing_data_formatter.HeaderRelatedData,
) error {
	var bPSupplierRecord *models.DataPlatformBusinessPartnerSupplierDatum
	var nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum
	var bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice
	var bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice
	var bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. ビジネスパートナ 得意先データ/仕入先データ の取得
		bPSupplierRecord, e = f.ExtractBusinessPartnerSupplierRecord(businessPartner, seller)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-2. OrderID
		nRLatestNumberRecord, e = f.ExtractNumberRangeLatestNumberRecord(serviceLabel, property)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		start := time.Now()
		// 2-1. ビジネスパートナマスタの取引先機能データの取得
		bPSupplierPartnerFunctionArray, e = f.ExtractBusinessPartnerSupplierPartnerFunctionArray(businessPartner, seller)
		if e != nil {
			err = e
			return
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 2-2. ビジネスパートナの一般データの取得
		bPGeneralArray, e = f.ExtractBusinessPartnerGeneralArray(sdc.Orders.HeaderPartner)
		if e != nil {
			err = e
			return
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 4-1. ビジネスパートナマスタの取引先プラントデータの取得
		bPSupplierPartnerPlantArray, e = f.ExtractBusinessPartnerSupplierPartnerPlantArray(businessPartner, seller, bPSupplierPartnerFunctionArray)
		if e != nil {
			err = e
			return
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	processingData = api_processing_data_formatter.ConvertToHeaderRelatedDataForBuyer(processingData, bPSupplierRecord, nRLatestNumberRecord, bPSupplierPartnerFunctionArray, bPSupplierPartnerPlantArray)
	f.l.Info(processingData)

	sdc.Orders = *api_output_formatter.ConvertToHeaderForBuyer(&sdc.Orders, bPSupplierRecord, nRLatestNumberRecord, bPSupplierPartnerFunctionArray, bPGeneralArray, bPSupplierPartnerPlantArray)

	return nil
}

func (f *SubFunction) CreateSdcForSeller(
	sdc *api_input_reader.SDC,
	businessPartner *int,
	serviceLabel string,
	property string,
	buyer *int,
	processingData *api_processing_data_formatter.HeaderRelatedData,
) error {
	var bPCustomerRecord *models.DataPlatformBusinessPartnerCustomerDatum
	var nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum
	var bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice
	var bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice
	var bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. ビジネスパートナ 得意先データ/仕入先データ の取得
		bPCustomerRecord, e = f.ExtractBusinessPartnerCustomerRecord(businessPartner, buyer)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-2. OrderID
		nRLatestNumberRecord, e = f.ExtractNumberRangeLatestNumberRecord(serviceLabel, property)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		start := time.Now()
		// 2-1. ビジネスパートナマスタの取引先機能データの取得
		bPCustomerPartnerFunctionArray, e = f.ExtractBusinessPartnerCustomerPartnerFunctionArray(businessPartner, buyer)
		if e != nil {
			err = e
			return
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 2-2. ビジネスパートナの一般データの取得
		bPGeneralArray, e = f.ExtractBusinessPartnerGeneralArray(sdc.Orders.HeaderPartner)
		if e != nil {
			err = e
			return
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 4-1. ビジネスパートナマスタの取引先プラントデータの取得
		bPCustomerPartnerPlantArray, e = f.ExtractBusinessPartnerCustomerPartnerPlantArray(businessPartner, buyer, bPCustomerPartnerFunctionArray)
		if e != nil {
			err = e
			return
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	processingData = api_processing_data_formatter.ConvertToHeaderRelatedDataForSeller(processingData, bPCustomerRecord, nRLatestNumberRecord, bPCustomerPartnerFunctionArray, bPCustomerPartnerPlantArray)
	f.l.Info(processingData)

	sdc.Orders = *api_output_formatter.ConvertToHeaderForSeller(&sdc.Orders, bPCustomerRecord, nRLatestNumberRecord, bPCustomerPartnerFunctionArray, bPGeneralArray, bPCustomerPartnerPlantArray)

	return nil
}
