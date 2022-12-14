package requests

type PaymentDueDate struct {
	InvoiceDocumentDate string `json:"InvoiceDocumentDate"`
	PaymentDueDate      string `json:"PaymentDueDate"`
}
