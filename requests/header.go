package requests

type Header struct {
	BuyerOrSeller            string
	OrderID                  *int `json:"OrderID"`
	OrderIDLatestNumber      *int
	Incoterms                string `json:"Incoterms"`
	PaymentTerms             string `json:"PaymentTerms"`
	PaymentMethod            string `json:"PaymentMethod"`
	BPAccountAssignmentGroup string `json:"BPAccountAssignmentGroup"`
	Buyer                    *int   `json:"Buyer"`
	Seller                   *int   `json:"Seller"`
	FieldNameWithNumberRange string
}
