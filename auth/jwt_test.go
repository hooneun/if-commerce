package auth_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hooneun/if-commerce/auth"
)

func TestJWTToken(t *testing.T) {
	req := new(auth.JWTTokenRequest)
	req.ID = 2
	req.Exp = time.Now().Add(time.Minute * 10)
	req.Key = "secret"

	token, err := auth.CreateJWTToken(*req)

	if err != nil {
		t.Errorf("Error, %s", err)
	}

	claim := new(auth.UserClaim)

	parseToken, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		t.Errorf("Error, %s", err)
	}

	if !parseToken.Valid {
		t.Errorf("Error, %s", err)
	}

	t.Log(claim.ID)
}
