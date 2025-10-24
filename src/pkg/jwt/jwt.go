package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	secretKey []byte
}

type JwtClaims struct {
	PhoneNumber string `json:"phone_number"`
	PlayerID    int64  `json:"player_id"`
		jwt.RegisteredClaims
}

func NewJwtService() *JWTService {
	key := os.Getenv("JWT_SECRET_KEY")
	return &JWTService{secretKey: []byte(key)}
}

func (jwtSvc JWTService) GenerateToken(phonenumber string, playerID int64) (string, error) {
	if len(jwtSvc.secretKey) == 0 {
		return "", errors.New("jwt secret key is empty")
	}

	claims := JwtClaims{
		PhoneNumber: phonenumber,
		PlayerID:    playerID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Hour)),
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
