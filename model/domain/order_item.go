package domain

import (
	"github.com/google/uuid"
	"time"
)

type OrderItem struct {
	OrderID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UnitPrice uint64    `gorm:"type:bigint;not null"`
	Quantity  uint32    `gorm:"type:int;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
