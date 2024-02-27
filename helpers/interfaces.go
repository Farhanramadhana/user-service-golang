package helpers

type HelperInterface interface {
	GenerateToken(userID int) (string, error)
	ValidateJWT(tokenString string) (*Claims, error)
}
