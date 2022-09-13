package server

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: os.Getenv("JWT_KEY"),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/auth/login" || c.Path() == "/auth/signup" {
				return true
			}

			return false
		},
	}
}
