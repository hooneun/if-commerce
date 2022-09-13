package auth_test

import (
	"testing"
	"time"

	"github.com/hooneun/if-commerce/auth"
)

func TestMakeJWTToken(t *testing.T) {
	req := auth.JWTTokenRequest{
		UniqueID: 1,
		Exp:      time.Now().Add(time.Minute * 1).Unix(),
		Key:      "secret",
	}

	_, err := auth.MakeJWTToken(req)

	if err != nil {
		t.Errorf("%s", err)
	}
}
