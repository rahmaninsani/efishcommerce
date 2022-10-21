package domain

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []Product `gorm:"many2many:product_categories;"`
}
