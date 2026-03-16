package users

import (
	"net/http"
	"encoding/json"
	"go_learn/http/config"
	"go_learn/http/helpers"
	"go_learn/http/models"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		helper.ResponseWrite(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	query := "INSERT INTO users(email, password, username) VALUES($1, $2, $3)"

	_, err = postgre.DB.Exec(query, user.Email, user.Password, user.Username)

	if err != nil {
		// helper.ResponseWrite(w, http.StatusInternalServerError, "Failed Create Users")
		http.Error(w, err.Error(), 500)
		return
	}

	// w.WriteHeader(http.StatusCreated)

	helper.ResponseWrite(w, http.StatusOK, "User Created")
}