package profile

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go_learn/http/helpers"
	"go_learn/http/config"
	"go_learn/http/models"

)

var DB *sql.DB

func GetProfile( w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(helper.UserIDKey)

	row := postgre.DB.QueryRow(`
		SELECT id, email, username
		FROM users
		WHERE id = $1
	`, userID)

	var u models.User

	err := row.Scan(&u.ID, &u.Email, &u.Username)

	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]any{
		"id" : u.ID,
		"email" : u.Email,
		"username" : u.Username,
	}

	json.NewEncoder(w).Encode(response)

}