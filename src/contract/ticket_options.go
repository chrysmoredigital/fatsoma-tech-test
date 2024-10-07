package contract

import "github.com/google/uuid"

type GetTicketOptionResponse struct {
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Allocation uint32    `json:"allocation"`
	ID         uuid.UUID `json:"id"`
}

type CreateTicketOptionRequest struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation uint32 `json:"allocation"`
}

type CreateTicketOptionResponse struct {
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Allocation uint32    `json:"allocation"`
	ID         uuid.UUID `json:"id"`
}
