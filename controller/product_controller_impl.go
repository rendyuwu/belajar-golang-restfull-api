package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rendyuwu/belajar-golang-restfull-api/helper"
	"github.com/rendyuwu/belajar-golang-restfull-api/model/web"
	"github.com/rendyuwu/belajar-golang-restfull-api/service"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(r, &productCreateRequest)

	productResponse := controller.ProductService.Create(r.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(r, &productUpdateRequest)

	id, err := strconv.Atoi(p.ByName("productId"))
	helper.PanicIfError(err)
	productUpdateRequest.Id = id

	productResponse := controller.ProductService.Update(r.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("productId"))
	helper.PanicIfError(err)

	controller.ProductService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("productId"))
	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productResponses := controller.ProductService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
