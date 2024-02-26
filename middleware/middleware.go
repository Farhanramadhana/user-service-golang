package middleware

import (
	"net/http"
	"strings"
	"userService/helpers"

	"github.com/labstack/echo/v4"
)

type Middleware struct {
	helpers helpers.HelperInterface
}

type MiddlewareInterface interface {
	AuthenticateMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

func NewMiddleware(helpers helpers.HelperInterface) MiddlewareInterface {
	return &Middleware{helpers}
}

// AuthenticateMiddleware validates the JWT token in the Authorization header
func (m *Middleware) AuthenticateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
		claims, err := m.helpers.ValidateJWT(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Invalid token")
		}

		c.Set("user_id", claims.Subject)
		return next(c)
	}
}
