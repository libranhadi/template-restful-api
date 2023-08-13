package controller
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"template-restful-api/service"
	"template-restful-api/helper"
	"template-restful-api/model/web"
	"encoding/json"
	"strconv"
)


type CategorControllerImpl struct {
	CategoryService service.CategoryService
}

func (c *CategorControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	jsonDecode := json.NewDecoder(r.Body);
	categoryCreateRequest := web.CategoryCreateRequest{}

	errDecode := jsonDecode.Decode(&categoryCreateRequest)
	helper.PanicIfError(errDecode)

	categoryResponse := c.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : categoryResponse,
	}

	w.Header().Add("Content-Type", "application-json");
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (c *CategorControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	jsonDecode := json.NewDecoder(r.Body);
	categoryRequest := web.CategoryUpdateRequest{}

	errDecode := jsonDecode.Decode(&categoryRequest)
	helper.PanicIfError(errDecode)

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
	
	w.Header().Add("Content-Type", "application-json");
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&webResponse)
	helper.PanicIfError(err)
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
	
	w.Header().Add("Content-Type", "application-json");
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (c *CategorControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := c.CategoryService.Get(r.Context())
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : categoryResponses,
	}
	
	w.Header().Add("Content-Type", "application-json")
	encode := json.NewEncoder(w)
	err := encode.Encode(&webResponse)
	helper.PanicIfError(err)
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

	w.Header().Add("Content-Type", "application-json")
	encode := json.NewEncoder(w)
	err := encode.Encode(&webResponse)
	helper.PanicIfError(err)
}
