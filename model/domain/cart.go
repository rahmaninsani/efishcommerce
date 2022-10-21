package domain

import (
	"github.com/google/uuid"
	"time"
)

type Cart struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CartItems []CartItem
}
