package handler

import (
	"net/http"

	"github.com/hooneun/if-commerce/auth"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	login := new(auth.LoginRequest)
	if err := c.Bind(login); err != nil {
		return err
	}

	user, err := auth.Login(*login, *h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) Signup(c echo.Context) error {
	return c.JSON(http.StatusOK, "Signup")
}
