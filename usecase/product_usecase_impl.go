package usecase

import (
	"efishcommerce/helper"
	"efishcommerce/model/web"
	"efishcommerce/repository"
)

type ProductUseCaseImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductUseCase(productRepository repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{ProductRepository: productRepository}
}

func (useCase ProductUseCaseImpl) FindBySlug(productSlug string) (web.ProductDetailResponse, error) {
	product, err := useCase.ProductRepository.FindBySlug(productSlug)
	if err != nil {
		return web.ProductDetailResponse{}, err
	}

	return helper.ToProductDetailResponse(product), nil
}

func (useCase ProductUseCaseImpl) FindAll() ([]web.ProductResponse, error) {
	products, err := useCase.ProductRepository.FindAll()
	if err != nil {
		return []web.ProductResponse{}, err
	}

	return helper.ToProductResponses(products), nil
}

func (useCase ProductUseCaseImpl) FindAllWithFilter(filters web.ProductFilterRequest) ([]web.ProductResponse, error) {
	products, err := useCase.ProductRepository.FindAllWithFilter(filters)
	if err != nil {
		return []web.ProductResponse{}, err
	}

	return helper.ToProductResponses(products), nil
}
