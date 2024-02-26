package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

// Claims structure to hold JWT claims
type Claims struct {
	jwt.StandardClaims
}

// AuthenticateMiddleware validates the JWT token in the Authorization header
func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		if tokenString == "" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Validate JWT
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "user_id", claims.Subject))
		next.ServeHTTP(w, r)
	})
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
	jwtSecret := os.Getenv("jwt_secret_key")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		fmt.Println(claims)
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}
