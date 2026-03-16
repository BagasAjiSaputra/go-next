package users

import (
	"net/http"
	"encoding/json"
	"go_learn/http/config"
	"go_learn/http/models"
	"go_learn/http/helpers"
	"github.com/google/uuid"
)

func DeleteUsers(w http.ResponseWriter, r*http.Request) {

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
		DELETE FROM users
		WHERE ID=$1
	`
	_, err = postgre.DB.Exec(query, u.ID)

	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseWrite(w, http.StatusOK, "User Deleted")

}