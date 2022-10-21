package usecase

import (
	"efishcommerce/helper"
	"efishcommerce/model/web"
	"efishcommerce/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseImpl struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{UserRepository: userRepository}
}

func (useCase UserUseCaseImpl) Login(request web.UserLoginRequest) (web.UserResponse, error) {
	user, err := useCase.UserRepository.FindByEmail(request.Email)
	if err != nil {
		return web.UserResponse{}, err
	}

	if user.Email == "" {
		return web.UserResponse{}, errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return web.UserResponse{}, errors.New("email or password is wrong")
	}

	return helper.ToUserResponse(user), nil
}

func (useCase UserUseCaseImpl) FindByEmail(email string) (web.UserResponse, error) {
	user, err := useCase.UserRepository.FindByEmail(email)

	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(user), nil
}
