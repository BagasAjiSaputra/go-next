package product

import (
	"encoding/json"
	"net/http"
	// "database/sql"
	"go_learn/http/helpers"
	"go_learn/http/config"	
	"go_learn/http/models"	
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {

	var request models.Product

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helper.ResponseWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	userID := r.Context().Value(helper.UserIDKey)

	_, err = postgre.DB.Exec(`
		INSERT INTO products (user_id, name, price)
		VALUES ($1, $2, $3)
	`, userID, request.Name, request.Price)

	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseWrite(w, http.StatusOK, "Product Created")

}