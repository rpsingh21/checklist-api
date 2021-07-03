package handler

import (
	"fmt"
	"net/http"
)

// JSONContentTypeMiddleware will add the json content type header for all endpoints
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("content-type", "application/json; charset=UTF-8")
		next.ServeHTTP(rw, r)
	})
}

// JWTAuthMiddleware token authentication
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		fmt.Println(token)
		// jwt.SigningMethodES256.Verify(token)
		// jwt.Parser
		// ctx.
		next.ServeHTTP(rw, r)
	})
}
