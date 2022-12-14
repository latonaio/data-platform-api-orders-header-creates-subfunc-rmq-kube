package requests

type OrderRegistrationType struct {
	ReferenceDocument     *int   `json:"ReferenceDocument"`
	ReferenceDocumentItem *int   `json:"ReferenceDocumentItem"`
	RegistrationType      string `json:"RegistrationType"`
}
