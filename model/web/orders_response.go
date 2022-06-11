package web

type OrdersResponse struct {
	Id          int `json:"id"`
	CustomerId  int `json:"customer-id"`
	TotalAmount int `json:"total-amount"`
}
