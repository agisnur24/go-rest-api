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

type OrdersServiceImpl struct {
	OrdersRepository repository.OrdersRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewOrdersService(ordersRepository repository.OrdersRepository, DB *sql.DB, validate *validator.Validate) OrdersService {
	return &OrdersServiceImpl{
		OrdersRepository: ordersRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *OrdersServiceImpl) Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := domain.Orders{
		CustomerId:  request.CustomerId,
		TotalAmount: request.TotalAmount,
	}

	order = service.OrdersRepository.Save(ctx, tx, order)
	return helper.ToOrdersResponse(order)
}

func (service *OrdersServiceImpl) Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrdersRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	order.CustomerId = request.CustomerId
	order.TotalAmount = request.TotalAmount

	order = service.OrdersRepository.Update(ctx, tx, order)

	return helper.ToOrdersResponse(order)
}

func (service *OrdersServiceImpl) Delete(ctx context.Context, ordersId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrdersRepository.FindById(ctx, tx, ordersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrdersRepository.Delete(ctx, tx, order)
}

func (service *OrdersServiceImpl) FindById(ctx context.Context, ordersId int) web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrdersRepository.FindById(ctx, tx, ordersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrdersResponse(order)
}

func (service *OrdersServiceImpl) FindAll(ctx context.Context) []web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := service.OrdersRepository.FindAll(ctx, tx)

	return helper.ToOrdersResponses(orders)
}
