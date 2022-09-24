package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTTokenRequest struct {
	ID  int64
	Exp time.Time
	Key string
}

type UserClaim struct {
	jwt.RegisteredClaims
	ID int64 `json:"id"`
}

type JWTToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewRequest(id int64, exp time.Time, key string) *JWTTokenRequest {
	return &JWTTokenRequest{
		ID:  id,
		Exp: exp,
		Key: key,
	}
}

func CreateJWTToken(t JWTTokenRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(t.Exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		ID: t.ID,
	})

	tokenStr, err := token.SignedString([]byte(t.Key))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
