package handler

import (
	"github.com/labstack/echo/v4"
)

type CartHandler interface {
	Create(c echo.Context) error
	CheckoutAll(c echo.Context) error
	FindByUserId(c echo.Context) error
}
