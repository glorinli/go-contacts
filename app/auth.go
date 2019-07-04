package app

import (
	"net/http"
	u "lens/utils"
	"strings"
	"go-contacts/models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {

		// Endpoints that need no authentication
		notAuth := []string{"api/user/new", "api/user/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization")

		// Token is missing
		if tokenHeader == "" {
			sendInvalidTokenResponse("Missing auto token")
			return
		}

		split := strings.split(tokenHeader, " ")
		if len(split) != 2 {
			sendInvalidTokenResponse("Invalid auto token")
			return
		}

		tokenPart := split[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, err) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			sendInvalidTokenResponse("Invalid auto token")
			return
		}

		// Token is invalid
		if !token.Valid {
			sendInvalidTokenResponse("Token is not valid")
			return
		}

		// Auth ok
		fmt.Sprintf("Use: %", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func sendInvalidTokenResponse(w http.ResponseWriter, message string) {
	response = u.Message(false, message)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	u.Response(w, response)
}