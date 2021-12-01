package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	err = godotenv.Load()
	key = os.Getenv("SECRET_KEY")
)

type Service interface {
	GenerateTokenUser(UserID string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateTokenUser(UserID string) (string, error) {
	claim := jwt.MapClaims{
		"user_id": UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, nil
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
