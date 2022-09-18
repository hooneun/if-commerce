package server

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_ACCESS_KEY")),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/auth/login" || c.Path() == "/auth/signup" || c.Path() == "/" {
				return true
			}

			return false
		},
	}
}
