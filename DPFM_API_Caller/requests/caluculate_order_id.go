package requests

type CalculateOrderID struct {
	OrderIDLatestNumber *int
	OrderID             *int `json:"OrderID"`
}
