package handler

import (
	"net/http"

	"github.com/hooneun/if-commerce/auth"
	"github.com/hooneun/if-commerce/db"
	"github.com/hooneun/if-commerce/response"
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

	res := response.Ok(0, "", user)

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) Signup(c echo.Context) error {
	signup := new(db.CreateUserParams)
	if err := c.Bind(signup); err != nil {
		return err
	}

	err := auth.Signup(*signup, *h.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Signup")
}
