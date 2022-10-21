package v1

import (
	"efishcommerce/handler"
	"github.com/labstack/echo/v4"
)

func NewProductRouter(group *echo.Group, productHandler handler.ProductHandler) {
	products := group.Group("/products")

	products.GET("", productHandler.FindAll)
	products.GET("/:slug", productHandler.FindBySlug)
}
