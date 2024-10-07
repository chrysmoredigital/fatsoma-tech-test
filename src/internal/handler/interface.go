package handler

import (
	"github.com/chrysmoredigital/fatsoma-tech-test/internal/model"
	"github.com/google/uuid"
)

type ticketOptionService interface {
	ClientError(error) bool
	GetTicketOption(uuid.UUID) (*model.TicketOption, error)
	CreateTicketOption(model.TicketOption) (uuid.UUID, error)
	CreatePurchase(model.Purchase) error
}
