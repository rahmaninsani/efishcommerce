package usecase

import (
	"efishcommerce/model/web"
	"github.com/google/uuid"
)

type OrderUseCase interface {
	Create(request web.OrderCreateRequest) (web.OrderResponse, error)
	Update(request web.OrderUpdateRequest) (web.OrderResponse, error)
	FindByUserId(userId uuid.UUID) ([]web.OrderResponse, error)
}
