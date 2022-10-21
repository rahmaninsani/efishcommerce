package repository

import (
	"efishcommerce/model/domain"
)

type UserRepository interface {
	FindByEmail(email string) (domain.User, error)
}
