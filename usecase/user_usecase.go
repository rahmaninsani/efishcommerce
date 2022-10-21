package usecase

import (
	"efishcommerce/model/web"
)

type UserUseCase interface {
	Login(request web.UserLoginRequest) (web.UserResponse, error)
	FindByEmail(email string) (web.UserResponse, error)
}
