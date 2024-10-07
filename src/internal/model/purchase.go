package model

import (
	"github.com/google/uuid"
)

type Purchase struct {
	ID             uuid.UUID
	TicketOptionID uuid.UUID
	UserID         uuid.UUID
	Quantity       uint32
}
