package helper

import (
	"context"
	"net/http"

	// "encoding/json"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CtxKey string

const UserIDKey CtxKey = "userID"

func JWTVerif(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			ResponseWrite(w, http.StatusUnauthorized, "Missing Token")
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			ResponseWrite(w, http.StatusUnauthorized, "Invalid Token Format")
			return
		}

		tokenString := parts[1]

		// tokenString := strings.Replace(authHeader, "Bearer", "", 1)

		// Counter alg:none attack
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			ResponseWrite(w, http.StatusUnauthorized, err.Error())
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := claims["id"]

		ctx := context.WithValue(r.Context(), UserIDKey, userID)


		next(w, r.WithContext(ctx))

	}
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID uuid.UUID, email string) (string, error) {

	claims := jwt.MapClaims{
		"id" : userID,
		"email" : email,
		"exp" : time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)

}