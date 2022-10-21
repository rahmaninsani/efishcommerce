package repository

import (
	"efishcommerce/model/domain"
	"efishcommerce/model/web"
	"github.com/google/uuid"
)

type ProductRepository interface {
	FindById(productId uuid.UUID) (domain.Product, error)
	FindBySlug(productSlug string) (domain.Product, error)
	FindAll() ([]domain.Product, error)
	FindAllWithFilter(filters web.ProductFilterRequest) ([]domain.Product, error)
}
