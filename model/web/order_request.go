package web

import (
	"github.com/google/uuid"
)

type OrderCreateRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Code   string
}

type OrderUpdateRequest struct {
	UserID                 uuid.UUID `json:"user_id"`
	OrderCode              string    `json:"order_code"`
	ProofOfPaymentFileName string    `json:"proof_of_payment_file_name"`
}
