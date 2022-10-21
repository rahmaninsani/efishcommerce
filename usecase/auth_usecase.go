package usecase

import (
	"efishcommerce/model/web"
	"github.com/golang-jwt/jwt"
)

type AuthUseCase interface {
	GenerateToken(userResponse web.UserResponse) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}
