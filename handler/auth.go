package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

	token, err := getTokenByUniqueID(user.ID)
	if err != nil {
		return err
	}

	res := response.Ok(http.StatusOK, "", token)

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

func getTokenByUniqueID(uniqueID int64) (*auth.JWTToken, error) {
	jwt := new(auth.JWTToken)

	accessKey := os.Getenv("JWT_ACCESS_KEY")
	refreshKey := os.Getenv("JWT_REFRESH_KEY")

	accessMin := os.Getenv("JWT_EXP_MIN")
	access, err := strconv.ParseInt(accessMin, 10, 64)
	if err != nil {
		return jwt, err
	}
	access = time.Now().Add(time.Minute * time.Duration(access)).Unix()
	fmt.Println(access)

	refreshMin := os.Getenv("JWT_REFRESH_EXP_MIN")
	refresh, err := strconv.ParseInt(refreshMin, 10, 64)
	if err != nil {
		return jwt, err
	}
	refresh = time.Now().Add(time.Minute * time.Duration(refresh)).Unix()

	accessData := auth.NewRequest(uniqueID, access, accessKey)
	refreshData := auth.NewRequest(uniqueID, refresh, refreshKey)

	accessToken, err := auth.MakeJWTToken(*accessData)
	if err != nil {
		return jwt, err
	}
	refreshToken, err := auth.MakeJWTToken(*refreshData)
	if err != nil {
		return jwt, err
	}

	jwt.AccessToken = accessToken
	jwt.RefreshToken = refreshToken

	return jwt, nil
}
