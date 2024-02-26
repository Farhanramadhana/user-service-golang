package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Helpers struct{}

func (h *Helpers) GenerateToken(userID int) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%v", userID),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
