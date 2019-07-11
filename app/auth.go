package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	u "github.com/glorinli/go-contacts/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/glorinli/go-contacts/models"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Endpoints that need no authentication
		notAuth := []string{"api/user/new", "api/user/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		// Token is missing
		if tokenHeader == "" {
			sendInvalidTokenResponse(w, "Missing auto token")
			return
		}

		split := strings.Split(tokenHeader, " ")
		if len(split) != 2 {
			sendInvalidTokenResponse(w, "Invalid auto token")
			return
		}

		tokenPart := split[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			sendInvalidTokenResponse(w, "Invalid auto token")
			return
		}

		// Token is invalid
		if !token.Valid {
			sendInvalidTokenResponse(w, "Token is not valid")
			return
		}

		// Auth ok
		fmt.Sprintf("Use: %", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func sendInvalidTokenResponse(w http.ResponseWriter, message string) {
	response := u.Message(false, message)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	u.Respond(w, response)
}
