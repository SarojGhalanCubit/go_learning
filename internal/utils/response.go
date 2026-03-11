package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool `json:"success"`
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}


func WriteSuccess(w http.ResponseWriter, status int, message string, data interface{}){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	
	response := APIResponse{
		Success: true,
		Status: status,
		Message: message,
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}

func WriteError(w http.ResponseWriter, status int, message string, err interface{}) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	response := APIResponse{
		Success: false,
		Status: status,
		Error: err,
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
