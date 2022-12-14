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
	OrderRegistrationType     *OrderRegistrationType     `json:"OrderRegistrationType"`
	OrderReferenceType        *OrderReferenceType        `json:"OrderReferenceType"`
	MetaData                  *MetaData                  `json:"MetaData"`
	BuyerSellerDetection      *BuyerSellerDetection      `json:"BuyerSellerDetection"`
	HeaderBPCustomer          *HeaderBPCustomer          `json:"HeaderBPCustomer"`
	HeaderBPSupplier          *HeaderBPSupplier          `json:"HeaderBPSupplier"`
	HeaderBPCustomerSupplier  *HeaderBPCustomerSupplier  `json:"HeaderBPCustomerSupplier"`
	CalculateOrderID          *CalculateOrderID          `json:"CalculateOrderID"`
	PaymentTerms              *[]PaymentTerms            `json:"PaymentTerms"`
	InvoiceDocumentDate       *InvoiceDocumentDate       `json:"InvoiceDocumentDate"`
	PaymentDueDate            *PaymentDueDate            `json:"PaymentDueDate"`
	NetPaymentDays            *NetPaymentDays            `json:"NetPaymentDays"`
	OverallDocReferenceStatus *OverallDocReferenceStatus `json:"OverallDocReferenceStatus"`
	PriceDetnExchangeRate     *PriceDetnExchangeRate     `json:"PriceDetnExchangeRate"`
	AccountingExchangeRate    *AccountingExchangeRate    `json:"AccountingExchangeRate"`
	TransactionCurrency       *TransactionCurrency       `json:"TransactionCurrency"`
	HeaderPartnerFunction     *[]HeaderPartnerFunction   `json:"HeaderPartnerFunction"`
	HeaderPartnerBPGeneral    *[]HeaderPartnerBPGeneral  `json:"HeaderPartnerBPGeneral"`
	HeaderPartnerPlant        *[]HeaderPartnerPlant      `json:"HeaderPartnerPlant"`
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
	ServiceLabel             *string `json:"ServiceLabel"`
	FieldNameWithNumberRange *string `json:"FieldNameWithNumberRange"`
	NumberRangeFrom          *int    `json:"NumberRangeFrom"`
	NumberRangeTo            *int    `json:"NumberRangeTo"`
}

type OrderReferenceType struct {
	ServiceLabel *string `json:"ServiceLabel"`
}

type BuyerSellerDetection struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
	Buyer             *int   `json:"Buyer"`
	Seller            *int   `json:"Seller"`
	BuyerOrSeller     string `json:"BuyerOrSeller"`
}

// Header
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
	FieldNameWithNumberRange string
}

type CalculateOrderIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	OrderIDLatestNumber      *int   `json:"OrderIDLatestNumber"`
}

type CalculateOrderID struct {
	OrderIDLatestNumber *int `json:"OrderIDLatestNumber"`
	OrderID             *int `json:"OrderID"`
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

type PriceDetnExchangeRate struct {
	PriceDetnExchangeRate *float32 `json:"PriceDetnExchangeRate"`
}

type AccountingExchangeRate struct {
	AccountingExchangeRate *float32 `json:"AccountingExchangeRate"`
}

type TransactionCurrency struct {
	TransactionCurrency string `json:"TransactionCurrency"`
}

type HeaderPartnerFunctionKey struct {
	OrderID            *int `json:"OrderID"`
	BusinessPartnerID  *int `json:"business_partner"`
	CustomerOrSupplier *int `json:"CustomerOrSupplier"`
}

type HeaderPartnerFunction struct {
	OrderID           *int   `json:"OrderID"`
	BusinessPartnerID *int   `json:"business_partner"`
	PartnerCounter    *int   `json:"PartnerCounter"`
	PartnerFunction   string `json:"PartnerFunction"`
	BusinessPartner   *int   `json:"BusinessPartner"`
	DefaultPartner    *bool  `json:"DefaultPartner"`
}

type HeaderPartnerBPGeneral struct {
	OrderID                 *int   `json:"OrderID"`
	BusinessPartner         *int   `json:"BusinessPartner"`
	BusinessPartnerFullName string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string `json:"BusinessPartnerName"`
	Country                 string `json:"Country"`
	Language                string `json:"Language"`
	AddressID               *int   `json:"AddressID"`
}

type HeaderPartnerPlantKey struct {
	OrderID                        *int   `json:"OrderID"`
	BusinessPartnerID              *int   `json:"business_partner"`
	CustomerOrSupplier             *int   `json:"CustomerOrSupplier"`
	PartnerCounter                 *int   `json:"PartnerCounter"`
	PartnerFunction                string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int   `json:"PartnerFunctionBusinessPartner"`
}

type HeaderPartnerPlant struct {
	OrderID         *int   `json:"OrderID"`
	BusinessPartner *int   `json:"BusinessPartner"`
	PartnerFunction string `json:"PartnerFunction"`
	PlantCounter    *int   `json:"PlantCounter"`
	Plant           string `json:"Plant"`
	DefaultPlant    *bool  `json:"DefaultPlant"`
}
