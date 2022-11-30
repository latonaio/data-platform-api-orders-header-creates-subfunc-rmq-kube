package requests

type HeaderPartnerPlantKey struct {
	OrderID                        *int `json:"OrderID"`
	BusinessPartnerID              *int `json:"business_partner"`
	CustomerOrSupplier             *int
	PartnerCounter                 *int   `json:"PartnerCounter"`
	PartnerFunction                string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int   `json:"PartnerFunctionBusinessPartner"`
}
