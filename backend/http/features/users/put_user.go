package users

import (
	"net/http"
	"encoding/json"
	"go_learn/http/config"
	"go_learn/http/models"
	"go_learn/http/helpers"
	// "github.com/google/uuid"
	"fmt"
)

func UpdateUsers(w http.ResponseWriter, r *http.Request) {

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		helper.ResponseWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	userID := r.Context().Value(helper.UserIDKey)

	// if u.ID == (uuid.UUID{}) {
	// 	helper.ResponseWrite(w, http.StatusBadRequest, "Missing User ID")
	// 	return
	// }

	query := `
	UPDATE users
	SET 
		email=$1, 
		password=$2, 
		username=$3, 
		phone=$4, 
		location=$5, 
		updated_at=now()
	WHERE id=$6`

	_, err = postgre.DB.Exec(
		query, 
		u.Email, 
		u.Password, 
		u.Username, 
		u.Phone, 
		u.Location, 
		// u.Role, 
		userID,
	)

	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(userID)

	helper.ResponseWrite(w, http.StatusOK, "User Updated")

}