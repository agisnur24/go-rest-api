package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
	"go-rest-api/repository"
)

type OrderProductServiceImpl struct {
	OrderProductRepository repository.OrderProductRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewOrderProductService(odpRepository repository.OrderProductRepository, DB *sql.DB, validate *validator.Validate) OrderProductService {
	return &OrderProductServiceImpl{
		OrderProductRepository: odpRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *OrderProductServiceImpl) Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	odp := domain.OrderProduct{
		OrderId:   request.OrderId,
		ProductId: request.ProductId,
		Qty:       request.Qty,
		Price:     request.Price,
		Amount:    request.Amount,
	}

	odp = service.OrderProductRepository.Save(ctx, tx, odp)
	return helper.ToOrderProductResponse(odp)
}

func (service *OrderProductServiceImpl) Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	odp, err := service.OrderProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	odp.OrderId = request.OrderId
	odp.ProductId = request.ProductId
	odp.Qty = request.Qty
	odp.Price = request.Price
	odp.Amount = request.Amount

	odp = service.OrderProductRepository.Update(ctx, tx, odp)

	return helper.ToOrderProductResponse(odp)
}

func (service *OrderProductServiceImpl) Delete(ctx context.Context, odpId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	odp, err := service.OrderProductRepository.FindById(ctx, tx, odpId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrderProductRepository.Delete(ctx, tx, odp)
}

func (service *OrderProductServiceImpl) FindById(ctx context.Context, odpId int) web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	odp, err := service.OrderProductRepository.FindById(ctx, tx, odpId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderProductResponse(odp)
}

func (service *OrderProductServiceImpl) FindAll(ctx context.Context) []web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	odp := service.OrderProductRepository.FindAll(ctx, tx)

	return helper.ToOrderProductResponses(odp)
}
