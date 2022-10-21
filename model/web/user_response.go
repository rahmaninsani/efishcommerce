package web

import (
	"github.com/google/uuid"
	"time"
)

type UserResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	AvatarFileName string    `json:"avatar_file_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
