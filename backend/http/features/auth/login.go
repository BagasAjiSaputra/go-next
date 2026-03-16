package feature

import (
	"encoding/json"
	"net/http"
	"database/sql"

	// "go_learn/http/database"
	"go_learn/http/config"
	"go_learn/http/helpers"
	"go_learn/http/models"
	"fmt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Email 		string `json:"email"`
		Password 	string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helper.ResponseWrite(w, http.StatusBadRequest, "Invalid Request")
		return	
	}

	var u models.User

	err = postgre.DB.QueryRow(`
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`, request.Email).Scan(&u.ID, &u.Email, &u.Password)

	if err != nil {

		if err == sql.ErrNoRows {
			helper.ResponseWrite(w, http.StatusUnauthorized, "Invalid Credentials")
			return
		}

		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return

	}

	// Compare Password
	if u.Password != request.Password {
		helper.ResponseWrite(w, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	// Generate Jhewete
	token, err := helper.GenerateToken(u.ID, u.Email)

	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, "Server Error 2")
		return
	}

	response := map[string]string {
		"token" : token,
	}


	// helper.ResponseWrite(w, http.StatusUnauthorized, "Invalid Credentials")
	fmt.Println("Login Success :", u.Email )
	json.NewEncoder(w).Encode(response)
}