package v1

import (
	"efishcommerce/handler"
	"efishcommerce/middleware"
	"github.com/labstack/echo/v4"
)

func NewOrderRouter(group *echo.Group, orderHandler handler.OrderHandler, authMiddleware *middleware.AuthMiddleware) {
	orders := group.Group("/orders")
	orders.Use(authMiddleware.Auth())

	orders.GET("", orderHandler.FindByUserId)
	orders.PUT("", orderHandler.Update)
}
