package product

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"go_learn/http/helpers"
	"go_learn/http/config"
	"github.com/google/uuid"
)

func CreateProduct( w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm( 10 << 20)
	if err != nil {
		helper.ResponseWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	price, _ := strconv.Atoi(priceStr)

	// User ID JWT
	userID := r.Context().Value(helper.UserIDKey)

	file, handler, err := r.FormFile("image")

	if err != nil {
		helper.ResponseWrite(w, http.StatusBadRequest, err.Error())
		return
	}

	defer file.Close()

	path := "./http/storage/uploads/products"

	ext := filepath.Ext(handler.Filename)

	filename := uuid.New().String() + ext


	// Create Folder
	os.MkdirAll(path, os.ModePerm)

	dst, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer dst.Close()

	io.Copy(dst, file)

	imageURL := filename

	_, err = postgre.DB.Exec(`
		INSERT INTO products (user_id, name, price, image_url)
		VALUES ($1, $2, $3, $4)
	`, userID, name, price, imageURL)

	if err != nil {
		helper.ResponseWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseWrite(w, http.StatusOK, "Product Created")

}