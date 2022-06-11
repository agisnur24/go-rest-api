package helper

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

func ToOrdersResponse(order domain.Orders) web.OrdersResponse {
	return web.OrdersResponse{
		Id:          order.Id,
		CustomerId:  order.CustomerId,
		TotalAmount: order.TotalAmount,
	}
}

func ToOrdersResponses(orders []domain.Orders) []web.OrdersResponse {
	var ordersResponses []web.OrdersResponse
	for _, order := range orders {
		ordersResponses = append(ordersResponses, ToOrdersResponse(order))
	}
	return ordersResponses
}
