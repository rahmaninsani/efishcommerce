package handler

import (
	"efishcommerce/helper"
	"efishcommerce/model/web"
	"efishcommerce/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CartHandlerImpl struct {
	CartUseCase  usecase.CartUseCase
	OrderUseCase usecase.OrderUseCase
}

func NewCartHandler(cartUseCase usecase.CartUseCase) CartHandler {
	return &CartHandlerImpl{CartUseCase: cartUseCase}
}

func (handler CartHandlerImpl) Create(c echo.Context) error {
	currentUser := c.Get("currentUser").(web.UserResponse)
	userId := currentUser.ID

	cartCreateRequest := web.CartCreateRequest{
		UserID: userId,
	}

	if err := c.Bind(&cartCreateRequest); err != nil {
		return err
	}

	cartResponse, err := handler.CartUseCase.Create(cartCreateRequest)
	helper.PanicIfError(err)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Success", cartResponse)
	return c.JSON(http.StatusCreated, webResponse)
}

func (handler CartHandlerImpl) CheckoutAll(c echo.Context) error {
	currentUser := c.Get("currentUser").(web.UserResponse)
	userId := currentUser.ID

	orderCreateRequest := web.OrderCreateRequest{
		UserID: userId,
		Code:   helper.GenerateOrderCode(),
	}

	if err := c.Bind(&orderCreateRequest); err != nil {
		return err
	}

	orderResponse, err := handler.CartUseCase.CheckoutAll(orderCreateRequest)
	helper.PanicIfError(err)

	webResponse := helper.WebResponse(http.StatusCreated, "OK", "Success", orderResponse)
	return c.JSON(http.StatusCreated, webResponse)
}

func (handler CartHandlerImpl) FindByUserId(c echo.Context) error {
	currentUser := c.Get("currentUser").(web.UserResponse)
	userId := currentUser.ID

	cartResponses, err := handler.CartUseCase.FindByUserId(userId)
	helper.PanicIfError(err)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Success", cartResponses)
	return c.JSON(http.StatusOK, webResponse)
}
