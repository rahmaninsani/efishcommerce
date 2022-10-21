package domain

import (
	"github.com/google/uuid"
	"time"
)

type ProductImage struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	ProductID uuid.UUID `gorm:"type:uuid"`
	FileName  string    `gorm:"type:varchar(45);not null"`
	IsPrimary bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
