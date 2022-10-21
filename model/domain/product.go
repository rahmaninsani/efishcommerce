package domain

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Price         uint64    `gorm:"type:bigint;not null"`
	Quantity      uint32    `gorm:"type:int;not null"`
	Description   string    `gorm:"type:text;not null"`
	Slug          string    `gorm:"type:varchar(200);not null;unique"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Categories    []Category `gorm:"many2many:product_categories;"`
	ProductImages []ProductImage
	CartItems     []CartItem
	OrderItems    []OrderItem
}
