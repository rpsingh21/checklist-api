package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/rpsingh21/checklist-api/utils"
)

// AuthContext hold session username
type AuthContext struct {
	Username        string
	Isauthenticated bool
}

// JSONContentTypeMiddleware will add the json content type header for all endpoints
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("content-type", "application/json; charset=UTF-8")
		next.ServeHTTP(rw, r)
	})
}

// JWTAuthMiddleware token authentication
func JWTAuthMiddleware(jw *utils.JWToken) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			authContext := AuthContext{Isauthenticated: false}
			bearerToken := r.Header.Get("Authorization")
			if strings.HasPrefix(bearerToken, "Bearer ") {
				tokenStr := strings.TrimPrefix(bearerToken, "Bearer ")
				fmt.Println(tokenStr)
				claims, err := jw.GetClaims(tokenStr)
				if err != nil {
					ErrorResponseWriter(rw, http.StatusUnauthorized, err)
					return
				}
				authContext.Username = claims.Username
				authContext.Isauthenticated = true
			}

			// Attached auth details with request context
			ctx := context.WithValue(r.Context(), AuthContext{}, authContext)
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}

// IsauthenticatedMiddleware Allow only Authenticated user
func IsauthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authContext := r.Context().Value(AuthContext{}).(AuthContext)
		if !authContext.Isauthenticated {
			ErrorResponseWriter(rw, http.StatusForbidden, jwt.NewValidationError("Login required", jwt.ValidationErrorIssuer))
			return
		}
		next.ServeHTTP(rw, r)
	})
}
