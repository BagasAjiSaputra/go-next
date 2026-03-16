package server

import (
	"fmt"
	"go_learn/http/config"
	"go_learn/http/features/auth"
	"go_learn/http/features/products"
	"go_learn/http/features/profiles"
	"go_learn/http/features/users"
	"go_learn/http/helpers"
	"github.com/joho/godotenv"

	// "go_learn/http/users"
	"net/http"

	"log"
)

func ServerOn() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgre.ConnectDB()

	http.HandleFunc("/", startServer)
	// http.HandleFunc("/users", db.GetUser)
	// http.HandleFunc("/users/create", db.CreateUser)
	// http.HandleFunc("/users/update", db.UpdateUser)
	// http.HandleFunc("/users/delete", db.DeleteUser)
	// http.HandleFunc("/user", helper.JWTVerif(helper.LogWare(UserHandler)))
	// http.HandleFunc("/user/", helper.JWTVerif(helper.LogWare(db.GetUserByID)))
	// http.HandleFunc("/register", db.CreateUser)


	http.HandleFunc("/login", feature.Login)
	http.HandleFunc("/register", users.CreateUsers)

	// PostgreSQL
	http.HandleFunc("/users", RouteUsers)

	// Get Profile
	http.HandleFunc("/profile", helper.JWTVerif(profile.GetProfile))

	// Create Product
	http.HandleFunc("/products", helper.JWTVerif(product.CreateProducts))

	http.HandleFunc("/product", helper.JWTVerif(product.CreateProduct))

	// Storage
	http.Handle("/uploads/",
		http.StripPrefix("/uploads/",
			http.FileServer(http.Dir("./http/storage/uploads"))))

	fmt.Println("Server Run On : 8080")
	// http.ListenAndServe(":1986", nil)
	
	log.Fatal(http.ListenAndServe(":8080", nil))

	
}

func startServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server Started")
}