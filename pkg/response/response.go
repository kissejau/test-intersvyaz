package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"error"`
}

type Success struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	w.WriteHeader(statusCode)
	data, _ := json.Marshal(Error{Message: message})
	w.Write(data)
}

func NewSuccessResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	w.WriteHeader(statusCode)
	data, _ := json.Marshal(Success{Message: message})
	w.Write(data)
}

func NewResponse(w http.ResponseWriter, r *http.Request, statusCode int, obj any) {
	w.WriteHeader(statusCode)
	data, _ := json.Marshal(obj)
	w.Write(data)
}