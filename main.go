package main

import (
	"efishcommerce/config"
	"efishcommerce/handler"
	customMiddleware "efishcommerce/middleware"
	"efishcommerce/repository"
	"efishcommerce/routes/api/v1"
	"efishcommerce/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitializeConstantValue()
	db := config.NewDB()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
		Output: e.Logger.Output(),
	}))
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Static("/images/avatars", "public/images/avatars")
	e.Static("/images/payments", "public/images/payments")

	// Repositories
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	cartRepository := repository.NewCartRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	orderItemRepository := repository.NewOrderItemRepository(db)

	// Use Cases/Services
	authUseCase := usecase.NewAuthUseCase()
	productUseCase := usecase.NewProductUseCase(productRepository)
	cartUseCase := usecase.NewCartUseCase(cartRepository, productRepository, orderRepository, orderItemRepository)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, productRepository)
	userUseCase := usecase.NewUserUseCase(userRepository)

	// Handlers/Controllers
	userHandler := handler.NewUserHandler(userUseCase, authUseCase)
	productHandler := handler.NewProductHandler(productUseCase)
	cartHandler := handler.NewCartHandler(cartUseCase)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	// Middlewares
	authMiddleware := customMiddleware.NewAuthMiddleware(authUseCase, userUseCase)

	// Routes
	apiV1 := e.Group("/api/v1")
	v1.NewUserRouter(apiV1, userHandler)
	v1.NewProductRouter(apiV1, productHandler)
	v1.NewCartRouter(apiV1, cartHandler, authMiddleware)
	v1.NewOrderRouter(apiV1, orderHandler, authMiddleware)

	address := fmt.Sprintf(":%s", config.APP_PORT)
	e.Logger.Fatal(e.Start(address))
}
