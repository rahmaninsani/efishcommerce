package usecase

import (
	"efishcommerce/config"
	"efishcommerce/model/web"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type jwtCustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	jwt.StandardClaims
}

type AuthUseCaseImpl struct {
}

func NewAuthUseCase() AuthUseCase {
	return &AuthUseCaseImpl{}
}

func (useCase *AuthUseCaseImpl) GenerateToken(userResponse web.UserResponse) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		UserID: userResponse.ID,
		Name:   userResponse.Name,
		Email:  userResponse.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	signedToken, err := token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (useCase *AuthUseCaseImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(config.SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
