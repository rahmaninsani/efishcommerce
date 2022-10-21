package web

import (
	"github.com/google/uuid"
)

type CartCreateRequest struct {
	UserID      uuid.UUID `json:"user_id"`
	ProductSlug string    `json:"product_slug"`
	Quantity    uint32    `json:"quantity"`
}
