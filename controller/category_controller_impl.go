package controller
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"template-restful-api/service"
	"template-restful-api/helper"
	"template-restful-api/model/web"
	"strconv"
)


type CategorControllerImpl struct {
	CategoryService service.CategoryService
}

func (c *CategorControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadRequestBody(r, categoryCreateRequest)

	categoryResponse := c.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : categoryResponse,
	}

	helper.WriteResponseJson(w, webResponse);
}

func (c *CategorControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryRequest := web.CategoryUpdateRequest{}
	helper.ReadRequestBody(r, categoryRequest)

	categoryId := params.ByName("id");
	id, errConv := strconv.Atoi(categoryId)
	helper.PanicIfError(errConv)

	categoryRequest.Id = id

	categoryResponse := c.CategoryService.Update(r.Context(), categoryRequest)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : categoryResponse,
	}

	helper.WriteResponseJson(w, webResponse);
}

func (c *CategorControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id");
	Id, errConv := strconv.Atoi(categoryId)
	helper.PanicIfError(errConv)

	c.CategoryService.Delete(r.Context(), Id)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
	}
	helper.WriteResponseJson(w, webResponse);
}

func (c *CategorControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := c.CategoryService.Get(r.Context())
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : categoryResponses,
	}

	helper.WriteResponseJson(w, webResponse);
}

func (c *CategorControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id");
	Id, errConv := strconv.Atoi(categoryId)
	helper.PanicIfError(errConv)

	categoryResponse := c.CategoryService.FindById(r.Context(), Id)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : categoryResponse,
	}

	helper.WriteResponseJson(w, webResponse);
}
