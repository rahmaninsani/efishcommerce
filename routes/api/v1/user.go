package v1

import (
	"efishcommerce/handler"
	"github.com/labstack/echo/v4"
)

func NewUserRouter(group *echo.Group, userHandler handler.UserHandler) {
	users := group.Group("/users")

	users.POST("/login", userHandler.Login)
}
