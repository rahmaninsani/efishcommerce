package usecase

import (
	"efishcommerce/helper"
	"efishcommerce/model/domain"
	"efishcommerce/model/web"
	"efishcommerce/repository"
	"github.com/google/uuid"
)

type CartUseCaseImpl struct {
	CartRepository      repository.CartRepository
	ProductRepository   repository.ProductRepository
	OrderRepository     repository.OrderRepository
	OrderItemRepository repository.OrderItemRepository
}

func NewCartUseCase(cartRepository repository.CartRepository, productRepository repository.ProductRepository,
	orderRepository repository.OrderRepository, orderItemRepository repository.OrderItemRepository) CartUseCase {
	return &CartUseCaseImpl{
		CartRepository:      cartRepository,
		ProductRepository:   productRepository,
		OrderRepository:     orderRepository,
		OrderItemRepository: orderItemRepository,
	}
}

func (useCase CartUseCaseImpl) Create(request web.CartCreateRequest) (web.CartResponse, error) {
	product, err := useCase.ProductRepository.FindBySlug(request.ProductSlug)
	helper.PanicIfError(err)

	cart, err := useCase.CartRepository.FindByUserId(request.UserID)
	helper.PanicIfError(err)

	cartItem := domain.CartItem{
		CartID:    cart.ID,
		ProductID: product.ID,
		Quantity:  request.Quantity,
	}

	cartItem, err = useCase.CartRepository.Save(cartItem)
	if err != nil {
		return web.CartResponse{}, err
	}

	return helper.ToCartResponse(cartItem, product), nil
}

func (useCase CartUseCaseImpl) CheckoutAll(request web.OrderCreateRequest) (web.OrderResponse, error) {
	cart, err := useCase.CartRepository.FindByUserId(request.UserID)
	if err != nil {
		return web.OrderResponse{}, err
	}

	order := domain.Order{
		UserID: request.UserID,
		Status: "pending",
		Code:   request.Code,
	}
	order, err = useCase.OrderRepository.Save(order)
	if err != nil {
		return web.OrderResponse{}, err
	}

	var orderItems []domain.OrderItem
	var orderedProducts []domain.Product
	for _, cartItem := range cart.CartItems {
		product, err := useCase.ProductRepository.FindById(cartItem.ProductID)
		helper.PanicIfError(err)

		orderItem := domain.OrderItem{
			OrderID:   order.ID,
			ProductID: product.ID,
			UnitPrice: product.Price,
			Quantity:  cartItem.Quantity,
		}
		orderItems = append(orderItems, orderItem)
		orderedProducts = append(orderedProducts, product)
	}

	orderItems, err = useCase.OrderItemRepository.Save(orderItems)
	helper.PanicIfError(err)

	order.OrderItems = orderItems
	return helper.ToOrderResponse(order, orderedProducts), nil
}

func (useCase CartUseCaseImpl) FindByUserId(userId uuid.UUID) ([]web.CartResponse, error) {
	cart, err := useCase.CartRepository.FindByUserId(userId)
	if err != nil {
		return []web.CartResponse{}, err
	}

	var products []domain.Product
	for _, cartItem := range cart.CartItems {
		product, err := useCase.ProductRepository.FindById(cartItem.ProductID)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return helper.ToCartResponses(cart, products), nil
}
