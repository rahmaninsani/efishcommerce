package v1

import (
	"efishcommerce/handler"
	"efishcommerce/middleware"
	"github.com/labstack/echo/v4"
)

func NewCartRouter(group *echo.Group, cartHandler handler.CartHandler, authMiddleware *middleware.AuthMiddleware) {
	cart := group.Group("/cart")
	cart.Use(authMiddleware.Auth())

	cart.GET("", cartHandler.FindByUserId)
	cart.POST("", cartHandler.Create)
	cart.POST("/orders", cartHandler.CheckoutAll)
}
