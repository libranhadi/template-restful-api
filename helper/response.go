package helper

import (
	"template-restful-api/model/web"

)

func CreateResponse(code int, status string, data interface{}) *web.WebResponse {
	webResponse := web.WebResponse{
		Code : code,
		Status : status,
		Data : data,
	}
	
	return &webResponse
}