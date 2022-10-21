package repository

import (
	"efishcommerce/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	var user domain.User

	if err := repository.DB.Debug().
		Where("email = (?)", email).
		First(&user).
		Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}
