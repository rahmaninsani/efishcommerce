package helper

import (
	"efishcommerce/model/domain"
	"efishcommerce/model/web"
)

func ToProductDetailResponse(product domain.Product) web.ProductDetailResponse {
	var images []string
	for _, image := range product.ProductImages {
		images = append(images, image.FileName)
	}

	productDetailResponse := web.ProductDetailResponse{
		Name:        product.Name,
		Images:      images,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Slug:        product.Slug,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	return productDetailResponse
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	var categories []string
	for _, category := range product.Categories {
		categories = append(categories, category.Name)
	}

	return web.ProductResponse{
		Name:       product.Name,
		Image:      product.ProductImages[0].FileName,
		Price:      product.Price,
		Quantity:   product.Quantity,
		Slug:       product.Slug,
		Categories: categories,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToCartResponse(cartItem domain.CartItem, product domain.Product) web.CartResponse {
	productResponse := ToProductResponse(product)

	return web.CartResponse{
		ProductDetail: productResponse,
		Quantity:      cartItem.Quantity,
		CreatedAt:     cartItem.CreatedAt,
		UpdatedAt:     cartItem.UpdatedAt,
	}
}

func ToCartResponses(cart domain.Cart, products []domain.Product) []web.CartResponse {
	var cartResponses []web.CartResponse
	for _, cartItem := range cart.CartItems {
		for _, product := range products {
			if cartItem.ProductID == product.ID {
				cartResponses = append(cartResponses, ToCartResponse(cartItem, product))
				break
			}
		}
	}

	return cartResponses
}

func ToOrderResponse(order domain.Order, orderedProducts []domain.Product) web.OrderResponse {
	var products []web.ProductOrderResponse
	var total uint64

	for _, orderItem := range order.OrderItems {
		for _, orderedProduct := range orderedProducts {
			if orderItem.ProductID == orderedProduct.ID {
				product := web.ProductOrderResponse{
					Name:      orderedProduct.Name,
					Image:     orderedProduct.ProductImages[0].FileName,
					Slug:      orderedProduct.Slug,
					UnitPrice: orderedProduct.Price,
					Quantity:  orderItem.Quantity,
					Subtotal:  orderedProduct.Price * uint64(orderItem.Quantity),
				}
				products = append(products, product)
				total += product.Subtotal
				break
			}
		}

	}

	return web.OrderResponse{
		OrderCode:              order.Code,
		Products:               products,
		Total:                  total,
		Status:                 order.Status,
		ProofOfPaymentFileName: order.ProofOfPaymentFileName,
		CreatedAt:              order.CreatedAt,
		UpdatedAt:              order.UpdatedAt,
	}
}

func ToOrderResponses(orders []domain.Order, orderedProducts [][]domain.Product) []web.OrderResponse {
	var orderResponses []web.OrderResponse

	for orderIndex, order := range orders {
		orderResponse := ToOrderResponse(order, orderedProducts[orderIndex])
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		AvatarFileName: user.AvatarFileName,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}
