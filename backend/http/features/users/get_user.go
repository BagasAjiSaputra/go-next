package users

import (
	"encoding/json"
	"net/http"
	// "go_learn/http/helpers"
	"go_learn/http/models"
	"go_learn/http/config"
	// "log"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := postgre.DB.Query("SELECT id, email, password, username, phone, location, updated_at FROM users")

	if err != nil {
		// helper.ResponseWrite(w, http.StatusInternalServerError, "Failed Query Select Users")
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var u models.User

		err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Username)

		if err != nil {
			// helper.ResponseWrite(w, http.StatusInternalServerError, "Failed Query Users")
			http.Error(w, err.Error(), 500)
			return
		}

		users = append(users, u)

	}

	// w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)

}