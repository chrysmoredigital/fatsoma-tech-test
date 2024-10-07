package api

import (
	"fmt"

	"github.com/chrysmoredigital/fatsoma-tech-test/internal/datastore/postgres"
	"github.com/chrysmoredigital/fatsoma-tech-test/internal/handler"
)

func (api *API) buildRoutes() error {
	var err error
	// deps
	api.pg, err = postgres.NewClient(postgres.NewDefaultConfig())
	if err != nil {
		return fmt.Errorf("failed to creat postgres client: %w", err)
	}

	// handlers
	ticketOptions, err := handler.NewTicketOptions(api.pg)
	if err != nil {
		return fmt.Errorf("failed to init ticket options handler: %w", err)
	}

	// routes
	api.r.Get("/ticket_options/{ticketOptionID}", ticketOptions.Get)
	api.r.Post("/ticket_options", ticketOptions.Create)
	api.r.Post("/ticket_options/{ticketOptionID}/purchases", ticketOptions.CreatePurchase)

	return nil
}
