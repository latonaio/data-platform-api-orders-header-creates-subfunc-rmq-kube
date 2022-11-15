package requests

type HeaderPartnerFunctionKey struct {
	OrderID            *int `json:"OrderID"`
	BusinessPartnerID  *int `json:"business_partner"`
	CustomerOrSupplier *int
}
