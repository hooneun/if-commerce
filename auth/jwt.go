package auth

import (
	"github.com/golang-jwt/jwt"
)

type JWTTokenRequest struct {
	UniqueID int64
	Exp      int64
	Key      string
}

type JWTToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewRequest(id int64, exp int64, key string) *JWTTokenRequest {
	return &JWTTokenRequest{
		UniqueID: id,
		Exp:      exp,
		Key:      key,
	}
}

func MakeJWTToken(t JWTTokenRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  t.UniqueID,
		"exp": t.Exp,
	})

	tokenStr, err := token.SignedString([]byte(t.Key))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
