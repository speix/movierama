package app

import "net/http"

/*
	ApiResponseService abstraction provides an interface for
	the Error response messages of the application to the client.
*/

type ApiResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type ApiResponseService interface {
	Respond(w http.ResponseWriter, code int, message string)
}
