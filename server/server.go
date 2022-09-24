package server

import (
	"net/http"

	"github.com/hooneun/if-commerce/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunAPI(address string) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(GetJWTConfig()))

	h, err := handler.NewHandler()
	if err != nil {
		return err
	}

	e.GET("/", h.HelloWorld)

	auth := e.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/signup", h.Signup)
	auth.GET("/me", h.Me)

	user := e.Group("/users")
	user.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET: /users")
	})
	user.POST("", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST: /users")
	})

	return e.Start(address)
}
