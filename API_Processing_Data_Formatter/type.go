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
	MetaData                  *MetaData                  `json:"MetaData"`
	Header                    *Header                    `json:"Header"`
	HeaderBPCustomerSupplier  *HeaderBPCustomerSupplier  `json:"HeaderBPCustomerSupplier"`
	BuyerSellerDetection      *BuyerSellerDetection      `json:"BuyerSellerDetection"`
	CalculateOrderIDKey       *CalculateOrderIDKey       `json:"CalculateOrderIDKey"`
	CalculateOrderIDQueryGets *CalculateOrderIDQueryGets `json:"CalculateOrderIDQueryGets"`
	CalculateOrderID          *CalculateOrderID          `json:"CalculateOrderID"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type BuyerSellerDetection struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
	Buyer             *int   `json:"Buyer"`
	Seller            *int   `json:"Seller"`
}

type Header struct {
	BuyerOrSeller            string
	OrderID                  *int `json:"OrderID"`
	OrderIDLatestNumber      *int
	Incoterms                string `json:"Incoterms"`
	PaymentTerms             string `json:"PaymentTerms"`
	PaymentMethod            string `json:"PaymentMethod"`
	BPAccountAssignmentGroup string `json:"BPAccountAssignmentGroup"`
	HeaderPartner            []HeaderPartner
	HeaderPartnerFunctionKey *HeaderPartnerFunctionKey `json:"HeaderPartnerFunctionKey"`
	HeaderPartnerFunction    *HeaderPartnerFunction    `json:"HeaderPartnerFunction"`
	HeaderPartnerBPGeneral   *HeaderPartnerBPGeneral   `json:"HeaderPartnerBPGeneral"`
}

type HeaderPartner struct {
	OrderID                 *int   `json:"OrderID"`
	BusinessPartnerID       *int   `json:"business_partner"`
	BusinessPartner         *int   `json:"BusinessPartner"`
	BusinessPartnerFullName string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string `json:"BusinessPartnerName"`
	Country                 string `json:"Country"`
	Language                string `json:"Language"`
	Currency                string `json:"Currency"`
	AddressID               *int   `json:"AddressID"`
	PartnerCounter          *int   `json:"PartnerCounter"`
	PartnerFunction         string `json:"PartnerFunction"`
	DefaultPartner          *bool  `json:"DefaultPartner"`
	HeaderPartnerPlant      []HeaderPartnerPlant
	HeaderPartnerPlantKey   *HeaderPartnerPlantKey `json:"HeaderPartnerPlantKey"`
}

type CalculateOrderIDKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string
}

type CalculateOrderIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string
	OrderIDLatestNumber      *int
}

type CalculateOrderID struct {
	OrderIDLatestNumber *int
	OrderID             *int `json:"OrderID"`
}

type HeaderBPCustomerSupplier struct {
	OrderID                  *int `json:"OrderID"`
	BusinessPartnerID        *int `json:"business_partner"`
	CustomerOrSupplier       *int
	Incoterms                string `json:"Incoterms"`
	PaymentTerms             string `json:"PaymentTerms"`
	PaymentMethod            string `json:"PaymentMethod"`
	BPAccountAssignmentGroup string `json:"BPAccountAssignmentGroup"`
}

type HeaderPartnerFunctionKey struct {
	OrderID            *int `json:"OrderID"`
	BusinessPartnerID  *int `json:"business_partner"`
	CustomerOrSupplier *int
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
	Currency                string `json:"Currency"`
	AddressID               *int   `json:"AddressID"`
}

type HeaderPartnerPlantKey struct {
	OrderID                        *int `json:"OrderID"`
	BusinessPartnerID              *int `json:"business_partner"`
	CustomerOrSupplier             *int
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

type HeaderPartnerRelatedData struct {
	PartnerFunction PartnerFunction `json:"PartnerFunction"`
	PartnerPlant    []PartnerPlant  `json:"PartnerPlant"`
}

type PartnerFunction struct {
	PartnerCounter *int  `json:"PartnerCounter"`
	DefaultPartner *bool `json:"DefaultPartner"`
}

type PartnerPlant struct {
	PlantCounter *int  `json:"PlantCounter"`
	DefaultPlant *bool `json:"DefaultPlant"`
}
