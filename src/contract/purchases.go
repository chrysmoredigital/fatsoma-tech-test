package contract

import "github.com/google/uuid"

type CreatePurchaseRequest struct {
	Quantity uint32    `json:"quantity"`
	UserID   uuid.UUID `json:"user_id"`
}
