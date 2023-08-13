package helper

import (
	"net/http"
	"encoding/json"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	jsonDecode := json.NewDecoder(r.Body);
	errDecode := jsonDecode.Decode(&result)
	PanicIfError(errDecode)
}

func WriteResponseJson(w http.ResponseWriter, response interface{})  {
	w.Header().Add("Content-Type", "application-json");
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&response)
	PanicIfError(err)
}