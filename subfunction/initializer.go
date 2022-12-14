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

func (f *SubFunction) OrderRegistrationType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.OrderRegistrationType, error) {
	var err error
	var orderRegistrationType *api_processing_data_formatter.OrderRegistrationType

	orderRegistrationType, err = psdc.ConvertToOrderRegistrationType(sdc)
	if err != nil {
		return nil, err
	}

	if orderRegistrationType.ReferenceDocument != nil {
		if orderRegistrationType.ReferenceDocumentItem != nil {
			orderRegistrationType.RegistrationType = "参照登録"
		} else {
			orderRegistrationType.RegistrationType = "直接登録"
		}
	} else {
		orderRegistrationType.RegistrationType = "直接登録"
	}
	return orderRegistrationType, nil
}

func (f *SubFunction) OrderReferenceType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.OrderReferenceType, error) {

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, NumberRangeFrom, NumberRangeTo
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_number_range_data`,
	)
	if err != nil {
		return nil, err
	}

	orderReferenceTypeQueryGets, err := psdc.ConvertToOrderReferenceTypeQueryGets(sdc, rows)
	if err != nil {
		return nil, err
	}

	data := &api_processing_data_formatter.OrderReferenceType{}

	for i := 0; i < len(*orderReferenceTypeQueryGets); i++ {
		if sdc.OrdersInputParameters.ReferenceDocument != nil && (*orderReferenceTypeQueryGets)[i].NumberRangeFrom != nil && (*orderReferenceTypeQueryGets)[i].NumberRangeTo != nil {
			if *sdc.OrdersInputParameters.ReferenceDocument >= *(*orderReferenceTypeQueryGets)[i].NumberRangeFrom && *sdc.OrdersInputParameters.ReferenceDocument <= *(*orderReferenceTypeQueryGets)[i].NumberRangeTo {
				data, err = psdc.ConvertToOrderReferenceType(sdc, &(*orderReferenceTypeQueryGets)[i])
				if err != nil {
					return nil, err
				}
				break
			}
		}
	}

	return data, err
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

	// 0-1 ReferenceDocumentおよびReferenceDocumentItemの値によるオーダー登録種別の判定
	psdc.OrderRegistrationType, e = f.OrderRegistrationType(sdc, psdc)
	if e != nil {
		err = e
		return err
	}

	//0-2. ReferenceDocumentの値によるオーダー参照先伝票種別の判定
	psdc.OrderReferenceType, e = f.OrderReferenceType(sdc, psdc)
	if e != nil {
		err = e
		return err
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. ビジネスパートナ 得意先データ/仕入先データ の取得
		psdc.HeaderBPCustomerSupplier, e = f.HeaderBPCustomerSupplier(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-3. InvoiceDocumentDate
		psdc.InvoiceDocumentDate, e = f.InvoiceDocumentDate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-4. PaymentDueDate
		psdc.PaymentDueDate, e = f.PaymentDueDate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-5. NetPaymentDays
		psdc.NetPaymentDays, e = f.NetPaymentDays(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		//1-7. OverallDocReferenceStatus
		psdc.OverallDocReferenceStatus, e = f.OverallDocReferenceStatus(psdc)
		if e != nil {
			err = e
			return
		}

		//1-9. PriceDetnExchangeRate
		psdc.PriceDetnExchangeRate, e = f.PriceDetnExchangeRate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		//1-10. AccountingExchangeRate
		psdc.AccountingExchangeRate, e = f.AccountingExchangeRate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 100-1. Currency
		e = f.Currency(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		//1-11. TransactionCurrency
		psdc.TransactionCurrency, e = f.TransactionCurrency(sdc, psdc)
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

	f.l.Info(psdc)
	osdc, err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
