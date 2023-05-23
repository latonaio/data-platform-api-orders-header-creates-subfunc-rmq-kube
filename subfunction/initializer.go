package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
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
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) OrderRegistrationType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.OrderRegistrationType {
	data := psdc.ConvertToOrderRegistrationType(sdc)

	return data
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

	orderReferenceTypeQueryGets, err := psdc.ConvertToOrderReferenceTypeQueryGets(rows)
	if err != nil {
		return nil, err
	}

	data := &api_processing_data_formatter.OrderReferenceType{}

	for i := 0; i < len(*orderReferenceTypeQueryGets); i++ {
		if sdc.OrdersInputParameters.ReferenceDocument != nil && (*orderReferenceTypeQueryGets)[i].NumberRangeFrom != nil && (*orderReferenceTypeQueryGets)[i].NumberRangeTo != nil {
			if *sdc.OrdersInputParameters.ReferenceDocument >= *(*orderReferenceTypeQueryGets)[i].NumberRangeFrom && *sdc.OrdersInputParameters.ReferenceDocument <= *(*orderReferenceTypeQueryGets)[i].NumberRangeTo {
				data = psdc.ConvertToOrderReferenceType(&(*orderReferenceTypeQueryGets)[i])
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

	psdc.MetaData = f.MetaData(sdc, psdc)

	// 0-1 ReferenceDocumentおよびReferenceDocumentItemの値によるオーダー登録種別の判定
	psdc.OrderRegistrationType = f.OrderRegistrationType(sdc, psdc)

	//0-2. ReferenceDocumentの値によるオーダー参照先伝票種別の判定
	psdc.OrderReferenceType, err = f.OrderReferenceType(sdc, psdc)
	if err != nil {
		return err
	}

	// 1-0. 入力ファイルのbusiness_partnerがBuyerであるかSellerであるかの判断
	psdc.BuyerSellerDetection, err = f.BuyerSellerDetection(sdc, psdc)
	if err != nil {
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

		// 2-2. ビジネスパートナの一般データの取得  // 2-1
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

		// 5-1-1. BPTaxClassification  // 2-2
		psdc.ItemBPTaxClassification, e = f.ItemBPTaxClassification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-1-2 ProductTaxClassification
		psdc.ItemProductTaxClassification, e = f.ItemProductTaxClassification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-1-3 TaxCode  // 5-1-1, 5-1-2
		psdc.TaxCode, e = f.TaxCode(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-8. PricingDate
		psdc.PricingDate = f.PricingDate(sdc, psdc)

		// 8-1. 価格マスタデータの取得(入力ファイルの[ConditionAmount]がnullである場合)  // 1-8
		psdc.PriceMaster, e = f.PriceMaster(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 8-2. 価格の計算(入力ファイルの[ConditionAmount]がnullである場合)  // 8-1
		psdc.ConditionAmount, e = f.ConditionAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-21. NetAmount  // 8-2
		psdc.NetAmount = f.NetAmount(sdc, psdc)

		// 10-1. TotalNetAmount  // 5-21
		psdc.TotalNetAmount, e = f.TotalNetAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-20. TaxRateの計算  // 5-1-3
		psdc.TaxRate, e = f.TaxRate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-22. TaxAmount  // 5-1-3, 5-20, 5-21
		psdc.TaxAmount, e = f.TaxAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 10-2. TotalTaxAmount  // 5-22
		psdc.TotalTaxAmount, e = f.TotalTaxAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 5-23. GrossAmount  // 5-21, 5-22
		psdc.GrossAmount, e = f.GrossAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 10-3. TotalGrossAmount  // 5-23
		psdc.TotalGrossAmount, e = f.TotalGrossAmount(sdc, psdc)
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
