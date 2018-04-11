package errors

import (
	"encoding/json"
	"net/http"

	"github.com/speix/movierama/app"
)

/*
	ApiResponseService implementation of Response endpoint.
*/
type ApiResponseService struct {
	app.ApiResponse
}

func (service *ApiResponseService) Respond(w http.ResponseWriter, code int, message string) {

	service.Error = true
	service.Message = message

	responseJson, _ := json.Marshal(service.ApiResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(responseJson)
}
