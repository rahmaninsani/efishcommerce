package web

import (
	"time"
)

type ProductOrderResponse struct {
	Name      string `json:"name"`
	Image     string `json:"image"`
	Slug      string `json:"slug"`
	UnitPrice uint64 `json:"unit_price"`
	Quantity  uint32 `json:"quantity"`
	Subtotal  uint64 `json:"subtotal"`
}

type OrderResponse struct {
	OrderCode              string                 `json:"order_code"`
	Products               []ProductOrderResponse `json:"products"`
	Total                  uint64                 `json:"total"`
	Status                 string                 `json:"status"`
	ProofOfPaymentFileName string                 `json:"proof_of_payment_file_name"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
}
