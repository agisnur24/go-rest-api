package app

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/controller"
	"go-rest-api/exception"
)

func NewRouter(category controller.CategoryController, customer controller.CustomerController, orderProduct controller.OrderProductController, orders controller.OrdersController, product controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", category.FindAll)
	router.GET("/api/categories/:categoryId", category.FindById)
	router.POST("/api/categories", category.Create)
	router.PUT("/api/categories/:categoryId", category.Update)
	router.DELETE("/api/categories/:categoryId", category.Delete)

	router.GET("/api/customers", customer.FindAll)
	router.GET("/api/customers/:customerId", customer.FindById)
	router.POST("/api/customers", customer.Create)
	router.PUT("/api/customers/:customerId", customer.Update)
	router.DELETE("/api/customers/:customerId", customer.Delete)

	router.GET("/api/odps", orderProduct.FindAll)
	router.GET("/api/odps/:odpId", orderProduct.FindById)
	router.POST("/api/odps", orderProduct.Create)
	router.PUT("/api/odps/:odpId", orderProduct.Update)
	router.DELETE("/api/odps/:odpId", orderProduct.Delete)

	router.GET("/api/orders", orders.FindAll)
	router.GET("/api/orders/:ordersId", orders.FindById)
	router.POST("/api/orders", orders.Create)
	router.PUT("/api/orders/:ordersId", orders.Update)
	router.DELETE("/api/orders/:ordersId", orders.Delete)

	router.GET("/api/products", product.FindAll)
	router.GET("/api/products/:productId", product.FindById)
	router.POST("/api/products", product.Create)
	router.PUT("/api/products/:productId", product.Update)
	router.DELETE("/api/products/:productId", product.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
