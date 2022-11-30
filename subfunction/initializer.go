package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"fmt"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

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
	psdc.MetaData = metaData

	buyerSellerDetection, err = psdc.ConvertToBuyerSellerDetection(sdc)
	if err != nil {
		return nil, err
	}

	// 1-0. 入力ファイルのbusiness_partnerがBuyerであるかSellerであるかの判断
	if *metaData.BusinessPartnerID == *buyerSellerDetection.Buyer && *metaData.BusinessPartnerID != *buyerSellerDetection.Seller {
		buyerSellerDetection.BuyerOrSeller = "Buyer"
		f.l.JsonParseOut(buyerSellerDetection.BuyerOrSeller)
	} else if *metaData.BusinessPartnerID != *buyerSellerDetection.Buyer && *metaData.BusinessPartnerID == *buyerSellerDetection.Seller {
		buyerSellerDetection.BuyerOrSeller = "Seller"
		f.l.JsonParseOut(buyerSellerDetection.BuyerOrSeller)
	} else {
		return nil, fmt.Errorf("business_partnerがBuyerまたはSellerと一致しません")
	}
	return buyerSellerDetection, nil
}

func (f *SubFunction) CalculateOrderID(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateOrderID, error) {
	metaData := psdc.MetaData
	dataKey, err := psdc.ConvertToCalculateOrderIDKey()
	if err != nil {
		return nil, err
	}

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateOrderIDQueryGets(sdc, rows)
	if err != nil {
		return nil, err
	}

	calculateOrderID := CalculateOrderID(*dataQueryGets.OrderIDLatestNumber)

	data, err := psdc.ConvertToCalculateOrderID(calculateOrderID)
	if err != nil {
		return nil, err
	}

	return data, err
}

func CalculateOrderID(latestNumber int) *int {
	res := latestNumber + 1
	return &res
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. ビジネスパートナ 得意先データ/仕入先データ の取得
		psdc.HeaderBPCustomerSupplier, e = f.HeaderBPCustomerSupplier(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-2. OrderID
		psdc.CalculateOrderID, e = f.CalculateOrderID(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 2-1. ビジネスパートナマスタの取引先機能データの取得
		psdc.HeaderPartnerFunction, e = f.HeaderPartnerFunction(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-2. ビジネスパートナの一般データの取得
		psdc.HeaderPartnerBPGeneral, e = f.HeaderPartnerBPGeneral(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 4-1. ビジネスパートナマスタの取引先プラントデータの取得
		psdc.HeaderPartnerPlant, e = f.HeaderPartnerPlant(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	osdc, err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
