package repository

import (
	"efishcommerce/model/domain"
	"gorm.io/gorm"
)

type OrderItemRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &OrderItemRepositoryImpl{DB: db}
}

func (repository OrderItemRepositoryImpl) Save(orderItems []domain.OrderItem) ([]domain.OrderItem, error) {
	if err := repository.DB.Debug().Create(&orderItems).Error; err != nil {
		return []domain.OrderItem{}, err
	}
	return orderItems, nil
}
