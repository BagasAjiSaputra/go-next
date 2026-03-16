package helper

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	Message string `json:"message"`
}



func ResponseWrite(w http.ResponseWriter, status int, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(Response{
		
		Message : message,
	})

}