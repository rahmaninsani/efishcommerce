package repository

import (
	"efishcommerce/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &CartRepositoryImpl{DB: db}
}

func (repository CartRepositoryImpl) Save(cartItem domain.CartItem) (domain.CartItem, error) {
	if err := repository.DB.Debug().Create(&cartItem).Error; err != nil {
		return domain.CartItem{}, err
	}
	return cartItem, nil
}

func (repository CartRepositoryImpl) Delete(cartItem domain.CartItem) error {
	if err := repository.DB.Debug().Delete(&domain.CartItem{}, cartItem.CartID).Error; err != nil {
		return err
	}
	return nil
}

func (repository CartRepositoryImpl) FindByUserId(userId uuid.UUID) (domain.Cart, error) {
	var cart domain.Cart

	if err := repository.DB.Debug().
		Preload("CartItems").
		Where("user_id = (?)", userId).
		Find(&cart).
		Error; err != nil {
		return domain.Cart{}, err
	}
	return cart, nil
}
