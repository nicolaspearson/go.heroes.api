package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"go-hero/models"
	u "go-hero/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtAuthentication : Authenticates a Jwt Auth Token
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// List of endpoints that don't require auth
		notAuth := []string{"/user/register", "/user/login", "/heroes"}
		// Current request path
		requestPath := r.URL.Path

		// Serve the request if auth is not required
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		// Grab the token from the header
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			// Token is missing, return error code 403 Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		// The token normally comes in the format `Bearer {token-body}`,
		// check if the retrieved token matches this format.
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid / malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1]
		// Grab the token
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			// Malformed token, return error code 403 Unauthorized
			response = u.Message(false, "Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			// Token is invalid
			response = u.Message(false, "Invalid auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// Proceed with the request and set the caller to the user retrieved from the parsed token.
		fmt.Sprintf("User %", tk.UserID)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		// Proceed in the middleware chain!
		next.ServeHTTP(w, r)
	})
}
