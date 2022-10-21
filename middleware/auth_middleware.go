package middleware

import (
	"efishcommerce/usecase"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"strings"
)

type AuthMiddleware struct {
	AuthUseCase usecase.AuthUseCase
	UserUseCase usecase.UserUseCase
}

func NewAuthMiddleware(authUseCase usecase.AuthUseCase, userUseCase usecase.UserUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		AuthUseCase: authUseCase,
		UserUseCase: userUseCase,
	}
}

func (middleware *AuthMiddleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
				return echo.ErrUnauthorized
			}

			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) < 2 {
				return echo.ErrUnauthorized
			}

			tokenString := arrayToken[1]

			token, err := middleware.AuthUseCase.ValidateToken(tokenString)

			if err != nil {
				return echo.ErrUnauthorized
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return echo.ErrUnauthorized
			}

			userEmail := claims["email"].(string)
			currentUser, err := middleware.UserUseCase.FindByEmail(userEmail)

			if err != nil {
				return echo.ErrUnauthorized
			}

			c.Set("currentUser", currentUser)
			if err = next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
