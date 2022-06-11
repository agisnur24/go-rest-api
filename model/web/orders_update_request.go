package web

type OrdersUpdateRequest struct {
	Id          int `validate:"required"`
	CustomerId  int `validate:"required" json:"customer-id"`
	TotalAmount int `validate:"required" json:"total-amount"`
}
