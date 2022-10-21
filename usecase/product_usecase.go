package usecase

import (
	"efishcommerce/model/web"
)

type ProductUseCase interface {
	FindBySlug(productSlug string) (web.ProductDetailResponse, error)
	FindAll() ([]web.ProductResponse, error)
	FindAllWithFilter(filters web.ProductFilterRequest) ([]web.ProductResponse, error)
}
