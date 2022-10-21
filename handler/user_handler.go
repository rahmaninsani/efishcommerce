package handler

import (
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Login(c echo.Context) error
}
