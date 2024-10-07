package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrysmoredigital/fatsoma-tech-test/contract"
	"github.com/chrysmoredigital/fatsoma-tech-test/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewTicketOptions(tos ticketOptionService) (*TicketOptions, error) {
	return &TicketOptions{
		svc: tos,
	}, nil
}

type TicketOptions struct {
	svc ticketOptionService
}

// Get is the HTTP handler for GET /ticket_options/{ticketOptionID}
func (h *TicketOptions) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "ticketOptionID"))
	if err != nil {
		respondError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("invalid ticket option id: %w", err),
		)
		return
	}

	to, err := h.svc.GetTicketOption(id)
	if err != nil {
		respondError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("failed to get ticket option: %w", err),
		)
		return
	}

	resp := contract.GetTicketOptionResponse{
		ID:         to.ID,
		Name:       to.Name,
		Desc:       to.Desc,
		Allocation: to.Allocation,
	}

	if err := respondJSON(w, http.StatusOK, resp); err != nil {
		log.Printf("failed to respond to request: %s", err)
	}
}

// Create is the HTTP handler for POST /ticket_options
func (h *TicketOptions) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	var req contract.CreateTicketOptionRequest

	if err = parseRequest(r, &req); err != nil {
		respondError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("parsing request: %w", err),
		)
		return
	}

	to := model.TicketOption{
		Name:       req.Name,
		Desc:       req.Desc,
		Allocation: req.Allocation,
	}

	to.ID, err = h.svc.CreateTicketOption(to)
	if err != nil {
		respondError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("failed to create ticket option: %w", err),
		)
		return
	}

	resp := contract.CreateTicketOptionResponse{
		ID:         to.ID,
		Name:       to.Name,
		Desc:       to.Desc,
		Allocation: to.Allocation,
	}

	if err := respondJSON(w, http.StatusOK, resp); err != nil {
		log.Printf("failed to respond to request: %s", err)
	}
}
