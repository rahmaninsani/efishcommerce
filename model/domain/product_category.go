package domain

import (
	"github.com/google/uuid"
	"time"
)

type ProductCategory struct {
	ProductID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	CategoryID uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
