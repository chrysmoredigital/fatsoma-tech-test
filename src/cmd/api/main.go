package main

import (
	"log"

	"github.com/chrysmoredigital/fatsoma-tech-test/internal/api"
)

func main() {
	api, err := api.New()
	if err != nil {
		log.Fatal("failed to init api: %w", err)
	}

	log.Fatal(
		api.ListenAndServe(),
	)
}
