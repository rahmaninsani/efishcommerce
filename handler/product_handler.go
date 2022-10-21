package handler

import (
	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	FindBySlug(c echo.Context) error
	FindAll(c echo.Context) error
}
