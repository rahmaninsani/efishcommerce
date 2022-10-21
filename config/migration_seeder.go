package config

import (
	"efishcommerce/helper"
	"efishcommerce/model/domain"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
)

func migrationSeeder(db *gorm.DB) error {
	models := []interface{}{
		&domain.Category{},
		&domain.Product{},
		&domain.ProductCategory{},
		&domain.ProductImage{},
		&domain.User{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
	}

	err := db.AutoMigrate(models...)
	helper.PanicIfError(err)
	log.Println("Migration Success")

	for _, model := range models[:6] {
		if db.Migrator().HasTable(model) {
			if err = db.First(model).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				switch model.(type) {
				case *domain.Category:
					err = batchInsertCategory(db)
					helper.PanicIfError(err)
				case *domain.Product:
					err = batchInsertProduct(db)
					helper.PanicIfError(err)
				case *domain.ProductCategory:
					err = batchInsertProductCategory(db)
					helper.PanicIfError(err)
				case *domain.ProductImage:
					err = batchInsertProductImage(db)
					helper.PanicIfError(err)
				case *domain.User:
					err = batchInsertUser(db)
					helper.PanicIfError(err)
				case *domain.Cart:
					err = batchInsertCart(db)
					helper.PanicIfError(err)
				}
				log.Println("Seeder Success for Model", reflect.TypeOf(model).Elem().Name())
			}
		}
	}

	return nil
}

func batchInsertCategory(db *gorm.DB) error {
	categories := []domain.Category{
		{Name: "Ikan"},
		{Name: "Ikan Konsumsi"},
		{Name: "Ikan Hias"},
		{Name: "Benih Ikan"},
		{Name: "Pakan"},
		{Name: "Suplemen"},
	}

	if err := db.Debug().Create(&categories).Error; err != nil {
		return err
	}
	return nil
}

func batchInsertProduct(db *gorm.DB) error {
	products := []domain.Product{
		{Name: "Benih Lele", Price: 50_000},
		{Name: "Pakan Lele Cap Menara", Price: 25_000},
		{Name: "Probiotik A", Price: 75_000},
		{Name: "Probiotik Nila B", Price: 10_000},
		{Name: "Pakan Nila", Price: 20_000},
		{Name: "Benih Nila Biasa", Price: 20_000},
		{Name: "Cupang", Price: 5_000},
		{Name: "Benih Nila Super", Price: 30_000},
		{Name: "Benih Cupang", Price: 10_000},
		{Name: "Probiotik B", Price: 10_000},
	}

	for index, product := range products {
		products[index].Quantity = 100
		products[index].Description = fmt.Sprintf("Deskripsi %s", strings.ToLower(product.Name))
		products[index].Slug = helper.GenerateSlug(product.Name)
	}

	if err := db.Debug().Create(&products).Error; err != nil {
		return err
	}
	return nil
}

func batchInsertProductCategory(db *gorm.DB) error {
	var products []domain.Product
	if err := db.Debug().Find(&products).Error; err != nil {
		return err
	}

	var categories []domain.Category
	if err := db.Debug().Find(&categories).Error; err != nil {
		return err
	}

	productItems := []domain.ProductCategory{
		{ProductID: products[0].ID, CategoryID: categories[0].ID},
		{ProductID: products[0].ID, CategoryID: categories[1].ID},
		{ProductID: products[0].ID, CategoryID: categories[3].ID},

		{ProductID: products[1].ID, CategoryID: categories[4].ID},

		{ProductID: products[2].ID, CategoryID: categories[5].ID},

		{ProductID: products[3].ID, CategoryID: categories[5].ID},

		{ProductID: products[4].ID, CategoryID: categories[4].ID},

		{ProductID: products[5].ID, CategoryID: categories[0].ID},
		{ProductID: products[5].ID, CategoryID: categories[1].ID},
		{ProductID: products[5].ID, CategoryID: categories[3].ID},

		{ProductID: products[6].ID, CategoryID: categories[0].ID},
		{ProductID: products[6].ID, CategoryID: categories[2].ID},

		{ProductID: products[7].ID, CategoryID: categories[0].ID},
		{ProductID: products[7].ID, CategoryID: categories[1].ID},
		{ProductID: products[7].ID, CategoryID: categories[3].ID},

		{ProductID: products[8].ID, CategoryID: categories[0].ID},
		{ProductID: products[8].ID, CategoryID: categories[2].ID},
		{ProductID: products[8].ID, CategoryID: categories[3].ID},

		{ProductID: products[9].ID, CategoryID: categories[5].ID},
	}
	if err := db.Debug().Create(&productItems).Error; err != nil {
		return err
	}
	return nil
}

func batchInsertProductImage(db *gorm.DB) error {
	var products []domain.Product
	if err := db.Debug().Find(&products).Error; err != nil {
		return err
	}

	var productImages []domain.ProductImage

	for _, product := range products {
		imageExtension := "jpg"

		for i := 0; i < 2; i++ {
			productImage := domain.ProductImage{
				ProductID: product.ID,
				FileName:  helper.GenerateFileName(imageExtension),
			}

			if i == 0 {
				productImage.IsPrimary = true
			}

			productImages = append(productImages, productImage)
		}
	}

	if err := db.Debug().Create(&productImages).Error; err != nil {
		return err
	}
	return nil
}

func batchInsertUser(db *gorm.DB) error {
	var users []domain.User
	userNames := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Helen", "Ivan", "John"}

	for _, userName := range userNames {
		user := domain.User{
			Name:         userName,
			Email:        fmt.Sprintf("%s@example.com", strings.ToLower(userName)),
			PasswordHash: helper.GenerateHashedPassword("123"),
		}

		users = append(users, user)
	}

	if err := db.Debug().Create(&users).Error; err != nil {
		return err
	}
	return nil
}

func batchInsertCart(db *gorm.DB) error {
	var users []domain.User
	if err := db.Debug().Find(&users).Error; err != nil {
		return err
	}

	var carts []domain.Cart

	for _, user := range users {
		cart := domain.Cart{
			UserID: user.ID,
		}

		carts = append(carts, cart)
	}

	if err := db.Debug().Create(&carts).Error; err != nil {
		return err
	}

	return nil
}
