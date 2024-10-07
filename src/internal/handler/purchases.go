package handler

import (
	"fmt"
	"net/http"

	"github.com/chrysmoredigital/fatsoma-tech-test/contract"
	"github.com/chrysmoredigital/fatsoma-tech-test/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// CreatePurchase is the HTTP handler for POST /ticket_options/{ticketOptionID}/purchases
func (h *TicketOptions) CreatePurchase(w http.ResponseWriter, r *http.Request) {
	var err error
	var req contract.CreatePurchaseRequest
	ticketOptionID := chi.URLParam(r, "ticketOptionID")

	if err = parseRequest(r, &req); err != nil {
		respondError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("parsing request: %w", err),
		)
		return
	}

	p := model.Purchase{
		UserID:   req.UserID,
		Quantity: req.Quantity,
	}

	p.TicketOptionID, err = uuid.Parse(ticketOptionID)
	if err != nil {
		respondError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("invalid ticket option id: %w", err),
		)
		return
	}

	err = h.svc.CreatePurchase(p)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if h.svc.ClientError(err) {
			statusCode = http.StatusBadRequest
		}
		respondError(
			w,
			statusCode,
			fmt.Errorf("failed to create purchase: %w", err),
		)
		return
	}

	respondStatus(w, http.StatusOK)
}
