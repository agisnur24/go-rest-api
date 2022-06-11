package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
	"go-rest-api/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	logrus.Info("Service Create Start")
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name:       request.Name,
		Price:      request.Price,
		CategoryId: request.CategoryId,
	}

	product = service.ProductRepository.Save(ctx, tx, product)
	logrus.Info("Service Create End")
	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	logrus.Info("Service Update Start")
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product := helper.ToProduct(productResponse)

	product.Name = request.Name
	product.Price = request.Price
	product.CategoryId = request.CategoryId

	product = service.ProductRepository.Update(ctx, tx, product)

	logrus.Info("Service Update End")
	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	logrus.Info("Service Delete Start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product := helper.ToProduct(productResponse)

	service.ProductRepository.Delete(ctx, tx, product)
	logrus.Info("Service Delete End")
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	logrus.Info("Service Find by Id Start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	logrus.Info("Service Find by Id End")
	return product
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	logrus.Info("Service Find All Start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	logrus.Info("Service Find All End")
	return products
}
