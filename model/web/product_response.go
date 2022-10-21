package web

import (
	"time"
)

type ProductResponse struct {
	Name       string    `json:"name"`
	Image      string    `json:"image"`
	Price      uint64    `json:"price"`
	Quantity   uint32    `json:"quantity"`
	Slug       string    `json:"slug"`
	Categories []string  `json:"categories"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductDetailResponse struct {
	Name        string    `json:"name"`
	Images      []string  `json:"image"`
	Price       uint64    `json:"price"`
	Quantity    uint32    `json:"quantity"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
