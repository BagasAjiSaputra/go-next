package server

import (
	"net/http"
	"go_learn/http/features/users"
)

func RouteUsers(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		users.GetUsers(w, r)

	// case http.MethodPost:
	// 	users.CreateUsers(w, r)
	
	case http.MethodPut:
		users.UpdateUsers(w, r)
	
	case http.MethodDelete:
		users.DeleteUsers(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}