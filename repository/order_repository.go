package repository

import (
	"efishcommerce/model/domain"
	"github.com/google/uuid"
)

type OrderRepository interface {
	Save(order domain.Order) (domain.Order, error)
	Update(order domain.Order) (domain.Order, error)
	FindByUserId(userId uuid.UUID) ([]domain.Order, error)
	FindByCode(orderCode string) (domain.Order, error)
}
