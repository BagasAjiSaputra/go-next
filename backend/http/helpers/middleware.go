package helper

import (
	"net/http"
	"fmt"
	"time"
)

func LogWare(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		fmt.Println("Request :", r.Method, r.URL.Path)
		next(w, r)
		fmt.Println("Done in :", time.Since(start))

	}

}