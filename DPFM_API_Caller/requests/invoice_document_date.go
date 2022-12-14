package requests

type InvoiceDocumentDate struct {
	RequestedDeliveryDate string `json:"RequestedDeliveryDate"`
	InvoiceDocumentDate   string `json:"InvoiceDocumentDate"`
}
