package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService struct {
	secretKey []byte
}

type JwtClaims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

func NewJwtService(secretKey []byte) *JWTService {
	return &JWTService{secretKey: secretKey}
}

func (jwtSvc JWTService) GenerateToken(phonenumber string) (string, error) {
	if len(jwtSvc.secretKey) == 0 {
		return "", errors.New("jwt secret key is empty")
	}

	claims := JwtClaims{
		PhoneNumber: phonenumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Subject:   "Access-Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString(jwtSvc.secretKey)

	if err != nil {
		return "", err
	}

	return newToken, nil
}
