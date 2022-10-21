package handler

import (
	"github.com/labstack/echo/v4"
)

type OrderHandler interface {
	Update(c echo.Context) error
	FindByUserId(c echo.Context) error
}
