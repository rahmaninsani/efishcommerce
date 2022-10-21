package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name           string    `gorm:"type:varchar(100);not null"`
	Email          string    `gorm:"type:varchar(100);not null; unique"`
	PasswordHash   string    `gorm:"type:varchar(60);not null"`
	AvatarFileName string    `gorm:"type:varchar(45);not null;default:'avatar_default.png'"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Cart           Cart
}
