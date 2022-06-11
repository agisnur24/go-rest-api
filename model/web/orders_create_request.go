package web

type OrdersCreateRequest struct {
	CustomerId  int `validate:"required" json:"customer-id"`
	TotalAmount int `validate:"required" json:"total-Amount"`
}
