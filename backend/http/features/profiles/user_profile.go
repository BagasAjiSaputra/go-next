package profile

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go_learn/http/helpers"
	"go_learn/http/config"
	"go_learn/http/models"
	"fmt"

)

var DB *sql.DB

func GetProfile( w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(helper.UserIDKey)

	row := postgre.DB.QueryRow(`
		SELECT id, email, password, username, phone, location, updated_at
		FROM users
		WHERE id = $1
	`, userID)

	var u models.User

	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Username, &u.Phone, &u.Location, &u.UpdatedAt)

	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]any{
		"id" : u.ID,
		"email" : u.Email,
		"password" : u.Password,
		"username" : u.Username,
		"phone" : u.Phone,
		"location" : u.Location,
		"updated_at" : u.UpdatedAt,
	}

	fmt.Println("User ID : ", userID)

	json.NewEncoder(w).Encode(response)

}