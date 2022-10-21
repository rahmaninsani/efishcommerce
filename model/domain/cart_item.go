package domain

import (
	"github.com/google/uuid"
	"time"
)

type CartItem struct {
	CartID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Quantity  uint32    `gorm:"type:int;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
