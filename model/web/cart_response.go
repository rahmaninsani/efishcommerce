package web

import "time"

type CartResponse struct {
	ProductDetail ProductResponse `json:"product_detail"`
	Quantity      uint32          `json:"quantity"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}
