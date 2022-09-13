package handler

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler() (*Handler, error) {
	db, err := sql.Open(os.Getenv("DATABASE"), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &Handler{DB: db}, nil
}

func (h *Handler) HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "IF Commerce")
}
