package repository

import (
	"efishcommerce/model/domain"
)

type OrderItemRepository interface {
	Save(orderItems []domain.OrderItem) ([]domain.OrderItem, error)
}
