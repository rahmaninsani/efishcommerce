package usecase

import (
	"efishcommerce/helper"
	"efishcommerce/model/domain"
	"efishcommerce/model/web"
	"efishcommerce/repository"
	"github.com/google/uuid"
)

type OrderUseCaseImpl struct {
	OrderRepository   repository.OrderRepository
	ProductRepository repository.ProductRepository
}

func NewOrderUseCase(orderRepository repository.OrderRepository, productRepository repository.ProductRepository) OrderUseCase {
	return &OrderUseCaseImpl{
		OrderRepository:   orderRepository,
		ProductRepository: productRepository,
	}
}

func (useCase OrderUseCaseImpl) Create(request web.OrderCreateRequest) (web.OrderResponse, error) {
	order := domain.Order{
		UserID: request.UserID,
	}

	order, err := useCase.OrderRepository.Save(order)
	if err != nil {
		return web.OrderResponse{}, err
	}

	return web.OrderResponse{}, nil
}

func (useCase OrderUseCaseImpl) Update(request web.OrderUpdateRequest) (web.OrderResponse, error) {
	order, err := useCase.OrderRepository.FindByCode(request.OrderCode)
	if err != nil {
		return web.OrderResponse{}, err
	}

	order.Status = "paid"
	order.ProofOfPaymentFileName = request.ProofOfPaymentFileName

	order, err = useCase.OrderRepository.Update(order)
	if err != nil {
		return web.OrderResponse{}, err
	}

	var orderedProducts []domain.Product
	for _, orderItem := range order.OrderItems {
		product, err := useCase.ProductRepository.FindById(orderItem.ProductID)
		helper.PanicIfError(err)
		orderedProducts = append(orderedProducts, product)
	}

	return helper.ToOrderResponse(order, orderedProducts), nil
}

func (useCase OrderUseCaseImpl) FindByUserId(userId uuid.UUID) ([]web.OrderResponse, error) {
	orders, err := useCase.OrderRepository.FindByUserId(userId)
	if err != nil {
		return []web.OrderResponse{}, err
	}

	var orderedProducts [][]domain.Product

	for _, order := range orders {
		var products []domain.Product
		for _, orderItem := range order.OrderItems {
			product, err := useCase.ProductRepository.FindById(orderItem.ProductID)
			helper.PanicIfError(err)
			products = append(products, product)
		}
		orderedProducts = append(orderedProducts, products)
	}

	return helper.ToOrderResponses(orders, orderedProducts), nil
}
