package repository

import (
	"efishcommerce/model/domain"
	"efishcommerce/model/web"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (repository ProductRepositoryImpl) FindById(productId uuid.UUID) (domain.Product, error) {
	var product domain.Product

	if err := repository.DB.Debug().
		Preload("Categories").
		Preload("ProductImages").
		Where("id = (?)", productId).
		First(&product).
		Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (repository ProductRepositoryImpl) FindBySlug(productSlug string) (domain.Product, error) {
	var product domain.Product
	if err := repository.DB.Debug().
		Preload("Categories").
		Preload("ProductImages").
		Where("slug = (?)", productSlug).
		First(&product).
		Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (repository ProductRepositoryImpl) FindAll() ([]domain.Product, error) {
	var products []domain.Product

	if err := repository.DB.Debug().
		Preload("Categories").
		Preload("ProductImages", "is_primary = (?)", true).
		Find(&products).
		Error; err != nil {
		return products, err
	}

	return products, nil
}

func (repository ProductRepositoryImpl) FindAllWithFilter(filters web.ProductFilterRequest) ([]domain.Product, error) {
	var products []domain.Product

	db := repository.DB.Debug().
		Joins("INNER JOIN product_categories ON product_categories.product_id = products.id").
		Joins("INNER JOIN categories ON product_categories.category_id = categories.id")

	if categories := filters.Categories; categories != nil {
		db.Where("categories.name IN (?)", categories)
	}

	switch {
	case filters.MinPrice > 0 && filters.MaxPrice > 0:
		db.Where("products.price BETWEEN (?) AND (?)", filters.MinPrice, filters.MaxPrice)
	case filters.MinPrice > 0:
		db.Where("products.price >= (?)", filters.MinPrice)
	case filters.MaxPrice > 0:
		db.Where("products.price <= (?)", filters.MaxPrice)
	}

	if err := db.Group("products.id").
		Preload("Categories").
		Preload("ProductImages", "is_primary = (?)", true).
		Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}
