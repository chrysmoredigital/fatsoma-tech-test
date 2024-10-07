package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrysmoredigital/fatsoma-tech-test/internal/datastore/postgres"
	"github.com/go-chi/chi/v5"
)

const (
	port = 3000
)

type API struct {
	r  chi.Router
	pg *postgres.Client
}

func New() (*API, error) {
	api := API{
		r: chi.NewMux(),
	}

	if err := api.buildRoutes(); err != nil {
		return nil, fmt.Errorf("failed to build routes: %w", err)
	}

	return &api, nil
}

func (api *API) ListenAndServe() (err error) {
	defer func() {
		err = api.gracefulShutdown()
	}()

	log.Printf("listening on port %d...", port)

	err = http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		api.r,
	)
	return
}

func (api *API) gracefulShutdown() error {
	log.Println("shutting down gracefully...")
	return api.pg.CloseDB()
}
