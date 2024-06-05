package jwt

import (
	"garasystem/pkg/util"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, secretKey string) (string, error) {
	claims := &MyClaims{
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(util.JWT_DURATION)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
