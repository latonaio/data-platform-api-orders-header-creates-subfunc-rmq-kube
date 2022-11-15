package requests

type HeaderPartnerBPGeneral struct {
	OrderID                        *int   `json:"OrderID"`
	BusinessPartner                *int   `json:"BusinessPartner"`
	PartnerFunctionBusinessPartner *int   `json:"PartnerFunctionBusinessPartner"`
	BusinessPartnerFullName        string `json:"BusinessPartnerFullName"`
	BusinessPartnerName            string `json:"BusinessPartnerName"`
	Country                        string `json:"Country"`
	Language                       string `json:"Language"`
	Currency                       string `json:"Currency"`
	AddressID                      *int   `json:"AddressID"`
}
