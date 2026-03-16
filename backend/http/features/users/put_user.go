package users

import (
	"net/http"
	"encoding/json"
	"go_learn/http/config"
	"go_learn/http/models"
	"go_learn/http/helpers"
	"github.com/google/uuid"
)

func UpdateUsers(w http.ResponseWriter, r *http.Request) {

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		helper.ResponseWrite(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if u.ID == (uuid.UUID{}) {
		helper.ResponseWrite(w, http.StatusBadRequest, "Missing User ID")
		return
	}

	query := `
	UPDATE users
	SET email=$1, password=$2, username=$3
	WHERE id=$4`

	_, err = postgre.DB.Exec(query, u.Email, u.Password, u.Username, u.ID)

	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseWrite(w, http.StatusOK, "User Updated")

}