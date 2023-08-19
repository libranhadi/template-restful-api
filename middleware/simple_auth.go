package middleware


import (
	"net/http"
	"template-restful-api/model/web"
	"template-restful-api/helper"
)

type SimpleAuth struct {
	Handler http.Handler
}


func NewSimpleAuthMiddleware(handler http.Handler) *SimpleAuth {
	return &SimpleAuth{Handler : handler}
}

func (auth *SimpleAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "SECRET" == r.Header.Get("X-API-KEY") {
		auth.Handler.ServeHTTP(w, r);
		return
	}
	resp := web.WebResponse{
		Code : http.StatusUnauthorized,
		Status : http.StatusText(http.StatusUnauthorized),
	}

	helper.WriteResponseJson(w, &resp);
	return
}