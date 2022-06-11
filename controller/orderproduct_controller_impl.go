package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"go-rest-api/service"
	"net/http"
	"strconv"
)

type OrderProductControllerImpl struct {
	OrderProductService service.OrderProductService
}

func NewOrderProductController(odpService service.OrderProductService) OrderProductController {
	return &OrderProductControllerImpl{
		OrderProductService: odpService,
	}
}

func (controller *OrderProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	odpCreateRequest := web.OrderProductCreateRequest{}
	helper.ReadFromRequestBody(request, &odpCreateRequest)

	odpResponse := controller.OrderProductService.Create(request.Context(), odpCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   odpResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	odpUpdateRequest := web.OrderProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &odpUpdateRequest)

	odpId := params.ByName("odpId")
	id, err := strconv.Atoi(odpId)
	helper.PanicIfError(err)

	odpUpdateRequest.Id = id

	odpResponse := controller.OrderProductService.Update(request.Context(), odpUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   odpResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	odpId := params.ByName("odpId")
	id, err := strconv.Atoi(odpId)
	helper.PanicIfError(err)

	controller.OrderProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	odpId := params.ByName("odpId")
	id, err := strconv.Atoi(odpId)
	helper.PanicIfError(err)

	odpResponse := controller.OrderProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   odpResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	odpResponse := controller.OrderProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   odpResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
