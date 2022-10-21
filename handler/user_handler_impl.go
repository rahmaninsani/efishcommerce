package handler

import (
	"efishcommerce/helper"
	"efishcommerce/model/web"
	"efishcommerce/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandlerImpl struct {
	UserUseCase usecase.UserUseCase
	AuthUseCase usecase.AuthUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase, authUseCase usecase.AuthUseCase) UserHandler {
	return &UserHandlerImpl{
		UserUseCase: userUseCase,
		AuthUseCase: authUseCase,
	}
}

func (handler UserHandlerImpl) Login(c echo.Context) error {
	var userLoginRequest web.UserLoginRequest
	webResponse := helper.WebResponse(http.StatusUnauthorized, "Error", "Unauthorized", nil)

	if err := c.Bind(&userLoginRequest); err != nil {
		return c.JSON(http.StatusUnauthorized, webResponse)
	}

	loggedInUser, err := handler.UserUseCase.Login(userLoginRequest)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, webResponse)
	}

	signedToken, err := handler.AuthUseCase.GenerateToken(loggedInUser)

	userLoginResponse := web.UserLoginResponse{
		Token: signedToken,
	}

	webResponse = helper.WebResponse(http.StatusOK, "OK", "Success", userLoginResponse)
	return c.JSON(http.StatusOK, webResponse)
}
