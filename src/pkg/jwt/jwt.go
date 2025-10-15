package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService struct {
	secretKey []byte
}

func NewJwtService(secretKey []byte) *JWTService {
	return &JWTService{secretKey: secretKey}
}

func (jwtSvc JWTService) GenerateToken() (string, error) {
	if len(jwtSvc.secretKey) == 0 {
		return "", errors.New("jwt secret key is empty")
	}

	claims := jwt.RegisteredClaims{
		Subject:   "Access-Token",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString(jwtSvc.secretKey)

	if err != nil {
		return "", err
	}

	return newToken, nil
}
