package helper

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

func ToOrderProductResponse(odp domain.OrderProduct) web.OrderProductResponse {
	return web.OrderProductResponse{
		Id:        odp.Id,
		OrderId:   odp.OrderId,
		ProductId: odp.ProductId,
		Qty:       odp.Qty,
		Price:     odp.Price,
		Amount:    odp.Amount,
	}
}

func ToOrderProductResponses(odps []domain.OrderProduct) []web.OrderProductResponse {
	var odpResponses []web.OrderProductResponse
	for _, odp := range odps {
		odpResponses = append(odpResponses, ToOrderProductResponse(odp))
	}
	return odpResponses
}
