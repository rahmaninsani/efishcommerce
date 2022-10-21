package repository

import (
	"efishcommerce/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{DB: db}
}

func (repository OrderRepositoryImpl) Save(order domain.Order) (domain.Order, error) {
	if err := repository.DB.Debug().Create(&order).Error; err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (repository OrderRepositoryImpl) Update(order domain.Order) (domain.Order, error) {
	if err := repository.DB.Debug().Save(&order).Error; err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (repository OrderRepositoryImpl) FindByUserId(userId uuid.UUID) ([]domain.Order, error) {
	var orders []domain.Order

	if err := repository.DB.Debug().
		Preload("OrderItems").
		Where("user_id = (?)", userId).
		Find(&orders).
		Error; err != nil {
		return []domain.Order{}, err
	}

	return orders, nil
}

func (repository OrderRepositoryImpl) FindByCode(orderCode string) (domain.Order, error) {
	var order domain.Order

	if err := repository.DB.Debug().
		Preload("OrderItems").
		Where("code = (?)", orderCode).
		Find(&order).
		Error; err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
