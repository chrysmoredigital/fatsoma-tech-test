package model

import (
	"github.com/google/uuid"
)

type TicketOption struct {
	Name       string
	Desc       string
	Allocation uint32
	ID         uuid.UUID
}
