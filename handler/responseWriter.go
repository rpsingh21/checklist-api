package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rpsingh21/checklist-api/model"
)

// ResponseWriter write Response to client
func ResponseWriter(rw http.ResponseWriter, statusCode int, message string, data interface{}) error {
	rw.WriteHeader(statusCode)

	if message != "" {
		data = model.NewResponse(statusCode, message, data)
	}
	httpResponse := json.NewEncoder(rw).Encode(data)
	return httpResponse
}

// ErrorResponseWriter return error meaage
func ErrorResponseWriter(rw http.ResponseWriter, statusCode int, err error) error {
	rw.WriteHeader(statusCode)
	message := err.Error()
	data := model.NewResponse(statusCode, message, nil)
	return json.NewEncoder(rw).Encode(data)
}
