package modularHTTP

import (
	. "ModularHTTPGo/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if user is authenticated
		if !IsAuthenticated(r) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// user is authenticated, call next handler
		next.ServeHTTP(w, r)
	})
}

func IsAuthenticated(r *http.Request) bool {
	result, err := VerifyJWT(r.Header.Get("Authorization"), []byte("my_secret_key"))
	if err != nil {
		return false
	}

	return result
}
