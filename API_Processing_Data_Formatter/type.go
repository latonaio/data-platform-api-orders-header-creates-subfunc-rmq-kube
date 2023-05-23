package api_processing_data_formatter

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	OrderRegistrationType        *OrderRegistrationType          `json:"OrderRegistrationType"`
	OrderReferenceType           *OrderReferenceType             `json:"OrderReferenceType"`
	MetaData                     *MetaData                       `json:"MetaData"`
	BuyerSellerDetection         *BuyerSellerDetection           `json:"BuyerSellerDetection"`
	HeaderBPCustomer             *HeaderBPCustomer               `json:"HeaderBPCustomer"`
	HeaderBPSupplier             *HeaderBPSupplier               `json:"HeaderBPSupplier"`
	HeaderBPCustomerSupplier     *HeaderBPCustomerSupplier       `json:"HeaderBPCustomerSupplier"`
	CalculateOrderID             *CalculateOrderID               `json:"CalculateOrderID"`
	PaymentTerms                 *[]PaymentTerms                 `json:"PaymentTerms"`
	InvoiceDocumentDate          *InvoiceDocumentDate            `json:"InvoiceDocumentDate"`
	PaymentDueDate               *PaymentDueDate                 `json:"PaymentDueDate"`
	NetPaymentDays               *NetPaymentDays                 `json:"NetPaymentDays"`
	OverallDocReferenceStatus    *OverallDocReferenceStatus      `json:"OverallDocReferenceStatus"`
	PricingDate                  *PricingDate                    `json:"PricingDate"`
	PriceDetnExchangeRate        *PriceDetnExchangeRate          `json:"PriceDetnExchangeRate"`
	AccountingExchangeRate       *AccountingExchangeRate         `json:"AccountingExchangeRate"`
	TransactionCurrency          *TransactionCurrency            `json:"TransactionCurrency"`
	TotalNetAmount               *TotalNetAmount                 `json:"TotalNetAmount"`
	TotalTaxAmount               *TotalTaxAmount                 `json:"TotalTaxAmount"`
	TotalGrossAmount             *TotalGrossAmount               `json:"TotalGrossAmount"`
	HeaderPartnerFunction        *[]HeaderPartnerFunction        `json:"HeaderPartnerFunction"`
	HeaderPartnerBPGeneral       *[]HeaderPartnerBPGeneral       `json:"HeaderPartnerBPGeneral"`
	HeaderPartnerPlant           *[]HeaderPartnerPlant           `json:"HeaderPartnerPlant"`
	ItemBPTaxClassification      *ItemBPTaxClassification        `json:"ItemBPTaxClassification"`
	ItemProductTaxClassification *[]ItemProductTaxClassification `json:"ItemProductTaxClassification"`
	TaxCode                      *[]TaxCode                      `json:"TaxCode"`
	TaxRate                      *[]TaxRate                      `json:"TaxRate"`
	NetAmount                    *[]NetAmount                    `json:"NetAmount"`
	TaxAmount                    *[]TaxAmount                    `json:"TaxAmount"`
	GrossAmount                  *[]GrossAmount                  `json:"GrossAmount"`
	PriceMaster                  *[]PriceMaster                  `json:"PriceMaster"`
	ConditionAmount              *[]ConditionAmount              `json:"ConditionAmount"`
}

// Initializer
type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type OrderRegistrationType struct {
	ReferenceDocument     *int   `json:"ReferenceDocument"`
	ReferenceDocumentItem *int   `json:"ReferenceDocumentItem"`
	RegistrationType      string `json:"RegistrationType"`
}

type OrderReferenceTypeQueryGets struct {
	ServiceLabel             string `json:"ServiceLabel"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	NumberRangeFrom          *int   `json:"NumberRangeFrom"`
	NumberRangeTo            *int   `json:"NumberRangeTo"`
}

type OrderReferenceType struct {
	ServiceLabel string `json:"ServiceLabel"`
}

// Header
type BuyerSellerDetection struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
	Buyer             *int   `json:"Buyer"`
	Seller            *int   `json:"Seller"`
	BuyerOrSeller     string `json:"BuyerOrSeller"`
}

type HeaderBPCustomer struct {
	OrderID                  *int    `json:"OrderID"`
	BusinessPartnerID        int     `json:"business_partner"`
	Customer                 int     `json:"Customer"`
	TransactionCurrency      *string `json:"TransactionCurrency"`
	Incoterms                *string `json:"Incoterms"`
	PaymentTerms             *string `json:"PaymentTerms"`
	PaymentMethod            *string `json:"PaymentMethod"`
	BPAccountAssignmentGroup *string `json:"BPAccountAssignmentGroup"`
}

type HeaderBPSupplier struct {
	OrderID                  *int    `json:"OrderID"`
	BusinessPartnerID        int     `json:"business_partner"`
	Supplier                 int     `json:"Supplier"`
	TransactionCurrency      *string `json:"TransactionCurrency"`
	Incoterms                *string `json:"Incoterms"`
	PaymentTerms             *string `json:"PaymentTerms"`
	PaymentMethod            *string `json:"PaymentMethod"`
	BPAccountAssignmentGroup *string `json:"BPAccountAssignmentGroup"`
}

type HeaderBPCustomerSupplier struct {
	OrderID                  *int    `json:"OrderID"`
	BusinessPartnerID        int     `json:"business_partner"`
	CustomerOrSupplier       int     `json:"CustomerOrSupplier"`
	TransactionCurrency      *string `json:"TransactionCurrency"`
	Incoterms                *string `json:"Incoterms"`
	PaymentTerms             *string `json:"PaymentTerms"`
	PaymentMethod            *string `json:"PaymentMethod"`
	BPAccountAssignmentGroup *string `json:"BPAccountAssignmentGroup"`
}

type CalculateOrderIDKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateOrderIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	OrderIDLatestNumber      *int   `json:"OrderIDLatestNumber"`
}

type CalculateOrderID struct {
	OrderIDLatestNumber *int `json:"OrderIDLatestNumber"`
	OrderID             int  `json:"OrderID"`
}

type PaymentTermsKey struct {
	PaymentTerms *string `json:"PaymentTerms"`
}

type PaymentTerms struct {
	PaymentTerms                string `json:"PaymentTerms"`
	BaseDate                    int    `json:"BaseDate"`
	BaseDateCalcAddMonth        *int   `json:"BaseDateCalcAddMonth"`
	BaseDateCalcFixedDate       *int   `json:"BaseDateCalcFixedDate"`
	PaymentDueDateCalcAddMonth  *int   `json:"PaymentDueDateCalcAddMonth"`
	PaymentDueDateCalcFixedDate *int   `json:"PaymentDueDateCalcFixedDate"`
}

type InvoiceDocumentDate struct {
	RequestedDeliveryDate string `json:"RequestedDeliveryDate"`
	InvoiceDocumentDate   string `json:"InvoiceDocumentDate"`
}

type PaymentDueDate struct {
	InvoiceDocumentDate string `json:"InvoiceDocumentDate"`
	PaymentDueDate      string `json:"PaymentDueDate"`
}

type NetPaymentDays struct {
	InvoiceDocumentDate string `json:"InvoiceDocumentDate"`
	PaymentDueDate      string `json:"PaymentDueDate"`
	NetPaymentDays      *int   `json:"NetPaymentDays"`
}

type OverallDocReferenceStatus struct {
	OverallDocReferenceStatus string `json:"OverallDocReferenceStatus"`
}

type PricingDate struct {
	PricingDate string `json:"PricingDate"`
}

type PriceDetnExchangeRate struct {
	PriceDetnExchangeRate *float32 `json:"PriceDetnExchangeRate"`
}

type AccountingExchangeRate struct {
	AccountingExchangeRate *float32 `json:"AccountingExchangeRate"`
}

type TransactionCurrency struct {
	TransactionCurrency string `json:"TransactionCurrency"`
}

type TotalNetAmount struct {
	TotalNetAmount float32 `json:"TotalNetAmount"`
}

type TotalTaxAmount struct {
	TotalTaxAmount float32 `json:"TotalTaxAmount"`
}

type TotalGrossAmount struct {
	TotalGrossAmount float32 `json:"TotalGrossAmount"`
}

// HeaderPartner
type HeaderPartnerFunctionKey struct {
	BusinessPartnerID  *int `json:"business_partner"`
	CustomerOrSupplier *int `json:"CustomerOrSupplier"`
}

type HeaderPartnerFunction struct {
	BusinessPartnerID int     `json:"business_partner"`
	PartnerCounter    int     `json:"PartnerCounter"`
	PartnerFunction   *string `json:"PartnerFunction"`
	BusinessPartner   *int    `json:"BusinessPartner"`
	DefaultPartner    *bool   `json:"DefaultPartner"`
}

type HeaderPartnerBPGeneral struct {
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string  `json:"BusinessPartnerName"`
	Country                 string  `json:"Country"`
	Language                string  `json:"Language"`
	AddressID               *int    `json:"AddressID"`
}

// HeaderPartnerPlant
type HeaderPartnerPlantKey struct {
	BusinessPartnerID              *int    `json:"business_partner"`
	CustomerOrSupplier             *int    `json:"CustomerOrSupplier"`
	PartnerCounter                 int     `json:"PartnerCounter"`
	PartnerFunction                *string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
}

type HeaderPartnerPlant struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	PartnerFunction               string  `json:"PartnerFunction"`
	PlantCounter                  int     `json:"PlantCounter"`
	Plant                         *string `json:"Plant"`
	DefaultPlant                  *bool   `json:"DefaultPlant"`
	DefaultStockConfirmationPlant *bool   `json:"DefaultStockConfirmationPlant"`
}

// Item
type ItemBPTaxClassificationKey struct {
	BusinessPartnerID  *int   `json:"business_partner"`
	CustomerOrSupplier *int   `json:"CustomerOrSupplier"`
	DepartureCountry   string `json:"DepartureCountry"`
}

type ItemBPTaxClassification struct {
	BusinessPartnerID   int     `json:"business_partner"`
	CustomerOrSupplier  int     `json:"CustomerOrSupplier"`
	DepartureCountry    string  `json:"DepartureCountry"`
	BPTaxClassification *string `json:"BPTaxClassification"`
}

type ItemProductTaxClassificationKey struct {
	Product           string `json:"Product"`
	BusinessPartnerID *int   `json:"business_partner"`
	Country           string `json:"Country"`
	TaxCategory       string `json:"TaxCategory"`
}

type ItemProductTaxClassification struct {
	Product                  string  `json:"Product"`
	BusinessPartnerID        int     `json:"business_partner"`
	Country                  string  `json:"Country"`
	TaxCategory              string  `json:"TaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}

type TaxCode struct {
	Product                  string  `json:"Product"`
	BPTaxClassification      *string `json:"BPTaxClassification"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
	OrderType                string  `json:"OrderType"`
	TaxCode                  string  `json:"TaxCode"`
}

type TaxRateKey struct {
	Country           string   `json:"Country"`
	TaxCode           []string `json:"TaxCode"`
	ValidityEndDate   string   `json:"ValidityEndDate"`
	ValidityStartDate string   `json:"ValidityStartDate"`
}

type TaxRate struct {
	Country           string   `json:"Country"`
	TaxCode           string   `json:"TaxCode"`
	ValidityEndDate   string   `json:"ValidityEndDate"`
	ValidityStartDate *string  `json:"ValidityStartDate"`
	TaxRate           *float32 `json:"TaxRate"`
}

type NetAmount struct {
	Product   string   `json:"Product"`
	NetAmount *float32 `json:"NetAmount"`
}

type TaxAmount struct {
	Product   string   `json:"Product"`
	TaxCode   string   `json:"TaxCode"`
	TaxRate   *float32 `json:"TaxRate"`
	NetAmount *float32 `json:"NetAmount"`
	TaxAmount *float32 `json:"TaxAmount"`
}

type GrossAmount struct {
	Product     string   `json:"Product"`
	NetAmount   *float32 `json:"NetAmount"`
	TaxAmount   *float32 `json:"TaxAmount"`
	GrossAmount *float32 `json:"GrossAmount"`
}

// ItemPricingElement
type PriceMasterKey struct {
	BusinessPartnerID          *int     `json:"business_partner"`
	Product                    []string `json:"Product"`
	CustomerOrSupplier         *int     `json:"CustomerOrSupplier"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
}

type PriceMaster struct {
	BusinessPartnerID          int      `json:"business_partner"`
	Product                    string   `json:"Product"`
	CustomerOrSupplier         *int     `json:"CustomerOrSupplier"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	ConditionType              string   `json:"ConditionType"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
}

type ConditionAmount struct {
	Product                    string   `json:"Product"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
