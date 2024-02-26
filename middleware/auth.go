package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Claims structure to hold JWT claims
type Claims struct {
	jwt.StandardClaims
}

// AuthenticateMiddleware validates the JWT token in the Authorization header
func AuthenticateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is missing")
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
		}

		// Validate JWT
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user_id", claims.Subject)
		return next(c)
	}
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

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
