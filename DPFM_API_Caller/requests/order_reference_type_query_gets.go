package requests

type OrderReferenceTypeQueryGets struct {
	ServiceLabel             *string `json:"ServiceLabel"`
	FieldNameWithNumberRange *string `json:"FieldNameWithNumberRange"`
	NumberRangeFrom          *int    `json:"NumberRangeFrom"`
	NumberRangeTo            *int    `json:"NumberRangeTo"`
}
