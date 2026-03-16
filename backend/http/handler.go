package server

import (
	"net/http"
	// "go_learn/http/users"
	"go_learn/http/features/users"
)


// func UserHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {
		
// 	case http.MethodGet:
// 		db.GetUser(w,r)

// 	// case http.MethodPost:
// 	// 	db.CreateUser(w,r)
		
// 	case http.MethodPut:
// 		db.UpdateUser(w,r)

// 	case http.MethodDelete:
// 		db.DeleteUser(w, r)

// 	default:
// 		http.Error(w, "Met	hod Not Allowed", http.StatusMethodNotAllowed)
// 	}

// }

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