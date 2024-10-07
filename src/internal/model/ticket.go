package model

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID             uuid.UUID
	PurchaseID     uuid.UUID
	TicketOptionID uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
