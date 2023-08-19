package exception


import (
	"net/http"
	"template-restful-api/model/web"
	"template-restful-api/helper"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{})  {
	if handlerNotFound(w, r, err) {
		return;
	}
	if hanlderInternalServerError(w, r, err) {
		return
	}

	if handlerBadRequest(w, r, err) {
		return
	}

}

func handlerBadRequest(w http.ResponseWriter, r *http.Request, err interface{}) bool{
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application-json");
		w.WriteHeader(http.StatusBadRequest);
	
		resp := web.WebResponse{
			Code : http.StatusBadRequest,
			Status : http.StatusText(http.StatusBadRequest),
			Data : exception.Error(),
		}
	
		helper.WriteResponseJson(w, &resp);
		return true
	}
	return false
}

func hanlderInternalServerError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(ErrorNotFound)
	if ok {
		w.Header().Set("Content-Type", "application-json");
		w.WriteHeader(http.StatusInternalServerError);
	
		resp := web.WebResponse{
			Code : http.StatusInternalServerError,
			Status : http.StatusText(http.StatusInternalServerError),
			Data : exception.Error,
		}
	
		helper.WriteResponseJson(w, &resp);
		return true
	}
	return false
}

func handlerNotFound(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(ErrorNotFound)
	if ok {
		w.Header().Set("Content-Type", "application-json");
		w.WriteHeader(http.StatusNotFound);

		resp := web.WebResponse{
			Code : http.StatusNotFound,
			Status : http.StatusText(http.StatusNotFound),
			Data : exception.Error,
		}
		helper.WriteResponseJson(w, &resp);
		return true;
	}
	return false;
}