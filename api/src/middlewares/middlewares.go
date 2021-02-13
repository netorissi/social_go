package middlewares

import (
	"api/src/auth"
	"api/src/utils"
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func UserAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidToken(r); err != nil {
			utils.AppError(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
