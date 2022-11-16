package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-orders-headers-creates-subfunc-rmq-kube/database"
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

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.MetaData, error) {
	var err error
	var metaData *api_processing_data_formatter.MetaData

	metaData, err = psdc.ConvertToMetaData(sdc)
	if err != nil {
		return nil, err
	}

	return metaData, nil
}

func (f *SubFunction) BuyerSellerDetection(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.BuyerSellerDetection, error) {
	var err error
	var buyerSellerDetection *api_processing_data_formatter.BuyerSellerDetection
	var metaData *api_processing_data_formatter.MetaData

	metaData, err = f.MetaData(sdc, psdc)
	if err != nil {
		return nil, err
	}

	buyerSellerDetection, err = psdc.ConvertToBuyerSellerDetection(sdc)
	if err != nil {
		return nil, err
	}

	// 1-0. 入力ファイルのbusiness_partnerがBuyerであるかSellerであるかの判断
	if *metaData.BusinessPartnerID == *buyerSellerDetection.Buyer && *metaData.BusinessPartnerID != *buyerSellerDetection.Seller {
		psdc.Header.BuyerOrSeller = "Buyer"
		f.l.JsonParseOut(psdc.Header.BuyerOrSeller)
	} else if *metaData.BusinessPartnerID != *buyerSellerDetection.Buyer && *metaData.BusinessPartnerID == *buyerSellerDetection.Seller {
		psdc.Header.BuyerOrSeller = "Seller"
		f.l.JsonParseOut(psdc.Header.BuyerOrSeller)
	} else {
		return nil, fmt.Errorf("business_partnerがBuyerまたはSellerと一致しません")
	}
	return buyerSellerDetection, nil
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	buyerSellerDetection *api_processing_data_formatter.BuyerSellerDetection,
) error {
	var headerBPCustomerSupplier *api_processing_data_formatter.HeaderBPCustomerSupplier
	var calculateOrderID *api_processing_data_formatter.CalculateOrderID
	var headerPartnerFunction *[]api_processing_data_formatter.HeaderPartnerFunction
	var headerPartnerBPGeneral *[]api_processing_data_formatter.HeaderPartnerBPGeneral
	var headerPartnerPlant *[]api_processing_data_formatter.HeaderPartnerPlant
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. ビジネスパートナ 得意先データ/仕入先データ の取得
		headerBPCustomerSupplier, e = f.HeaderBPCustomerSupplier(buyerSellerDetection, sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-2. OrderID
		calculateOrderID, e = f.CalculateOrderID(buyerSellerDetection, sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		start := time.Now()
		// 2-1. ビジネスパートナマスタの取引先機能データの取得
		headerPartnerFunction, e = f.HeaderPartnerFunction(buyerSellerDetection, sdc, psdc)
		if e != nil {
			err = e
			return
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 2-2. ビジネスパートナの一般データの取得
		headerPartnerBPGeneral, e = f.HeaderPartnerBPGeneral(headerPartnerFunction, sdc, psdc)
		if e != nil {
			err = e
			return
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 4-1. ビジネスパートナマスタの取引先プラントデータの取得
		headerPartnerPlant, e = f.HeaderPartnerPlant(buyerSellerDetection, headerPartnerFunction, sdc, psdc)
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

	sdc, err = f.SetValue(sdc, buyerSellerDetection, headerBPCustomerSupplier, calculateOrderID, headerPartnerFunction, headerPartnerBPGeneral, headerPartnerPlant)
	if err != nil {
		return err
	}

	return nil
}
