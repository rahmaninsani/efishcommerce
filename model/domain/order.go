package domain

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserID                 uuid.UUID `gorm:"type:uuid"`
	Status                 string    `gorm:"type:varchar(20);not null;default:'pending'"`
	Code                   string    `gorm:"type:varchar(23);not null;unique"`
	ProofOfPaymentFileName string    `gorm:"type:varchar(45);default:null"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
	OrderItems             []OrderItem
}
