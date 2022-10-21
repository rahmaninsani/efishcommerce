package usecase

import (
	"efishcommerce/model/web"
	"github.com/google/uuid"
)

type CartUseCase interface {
	Create(request web.CartCreateRequest) (web.CartResponse, error)
	CheckoutAll(request web.OrderCreateRequest) (web.OrderResponse, error)
	FindByUserId(userId uuid.UUID) ([]web.CartResponse, error)
}
