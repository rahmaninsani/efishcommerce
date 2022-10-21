package handler

import (
	"efishcommerce/helper"
	"efishcommerce/model/web"
	"efishcommerce/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductHandlerImpl struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductHandler(productUseCase usecase.ProductUseCase) ProductHandler {
	return &ProductHandlerImpl{
		ProductUseCase: productUseCase,
	}
}

func (handler ProductHandlerImpl) FindBySlug(c echo.Context) error {
	productSlug := c.Param("slug")
	productResponse, err := handler.ProductUseCase.FindBySlug(productSlug)
	helper.PanicIfError(err)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Success", productResponse)
	return c.JSON(http.StatusOK, webResponse)
}

func (handler ProductHandlerImpl) FindAll(c echo.Context) error {
	var productResponses []web.ProductResponse
	var err error

	if len(c.QueryParams()) > 0 {
		filters := web.ProductFilterRequest{}

		if categories, ok := c.QueryParams()["category"]; ok {
			filters.Categories = categories
		}

		if minPrice, err := strconv.ParseUint(c.QueryParam("min_price"), 10, 64); err == nil {
			filters.MinPrice = minPrice
		}

		if maxPrice, err := strconv.ParseUint(c.QueryParam("max_price"), 10, 64); err == nil {
			filters.MaxPrice = maxPrice
		}

		productResponses, err = handler.ProductUseCase.FindAllWithFilter(filters)
		helper.PanicIfError(err)
	} else {
		productResponses, err = handler.ProductUseCase.FindAll()
		helper.PanicIfError(err)
	}

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Success", productResponses)
	return c.JSON(http.StatusOK, webResponse)
}
