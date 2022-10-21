package repository

import (
	"efishcommerce/model/domain"
	"github.com/google/uuid"
)

type CartRepository interface {
	Save(cartItem domain.CartItem) (domain.CartItem, error)
	Delete(cartItem domain.CartItem) error
	FindByUserId(userId uuid.UUID) (domain.Cart, error)
}
