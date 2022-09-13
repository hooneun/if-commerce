package main

import (
	"log"
	"net/http"

	"github.com/hooneun/if-commerce/config"
	"github.com/hooneun/if-commerce/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(RunAPI(":1323"))
}

func RunAPI(address string) error {
	e := echo.New()
	e.Use(middleware.Logger())

	h, err := handler.NewHandler()
	if err != nil {
		return err
	}

	e.GET("/", h.HelloWorld)

	auth := e.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/signup", h.Signup)

	user := e.Group("/users")
	user.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET: /users")
	})
	user.POST("", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST: /users")
	})

	return e.Start(address)
}
