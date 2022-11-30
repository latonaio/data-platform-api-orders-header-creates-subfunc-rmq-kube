package requests

type HeaderPartnerPlant struct {
	OrderID         *int   `json:"OrderID"`
	BusinessPartner *int   `json:"BusinessPartner"`
	PartnerFunction string `json:"PartnerFunction"`
	PlantCounter    *int   `json:"PlantCounter"`
	Plant           string `json:"Plant"`
	DefaultPlant    *bool  `json:"DefaultPlant"`
}
