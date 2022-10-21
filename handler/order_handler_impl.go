package handler

import (
	"efishcommerce/helper"
	"efishcommerce/model/web"
	"efishcommerce/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strings"
)

type OrderHandlerImpl struct {
	OrderUseCase usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase usecase.OrderUseCase) OrderHandler {
	return &OrderHandlerImpl{OrderUseCase: orderUseCase}
}

func (handler OrderHandlerImpl) Update(c echo.Context) error {
	currentUser := c.Get("currentUser").(web.UserResponse)
	userId := currentUser.ID
	orderCode := c.FormValue("order_code")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	fileExtension := strings.Split(file.Filename, ".")
	fileNameWithExtension := helper.GenerateFileName(fileExtension[len(fileExtension)-1])

	path := fmt.Sprintf("public/images/payments/%s", fileNameWithExtension)

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	orderUpdateRequest := web.OrderUpdateRequest{
		UserID:                 userId,
		OrderCode:              orderCode,
		ProofOfPaymentFileName: fileNameWithExtension,
	}

	orderResponse, err := handler.OrderUseCase.Update(orderUpdateRequest)
	helper.PanicIfError(err)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Success", orderResponse)
	return c.JSON(http.StatusOK, webResponse)
}

func (handler OrderHandlerImpl) FindByUserId(c echo.Context) error {
	currentUser := c.Get("currentUser").(web.UserResponse)
	userId := currentUser.ID

	orderResponses, err := handler.OrderUseCase.FindByUserId(userId)
	helper.PanicIfError(err)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Success", orderResponses)
	return c.JSON(http.StatusOK, webResponse)
}
