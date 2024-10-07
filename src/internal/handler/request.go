package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* This file contains helper funcs for handling HTTP requests */

// parseRequest attempts to parse the incoming http request body into the
// destination as JSON. Destination should be pointer to struct.
func parseRequest(r *http.Request, dest interface{}) error {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(dest)
	if err != nil {
		return fmt.Errorf("decoding request body: %w", err)
	}

	return nil
}

// respondJSON responds to the HTTP request by attempting to marshal the response
// object as JSON
func respondJSON(w http.ResponseWriter, statusCode int, resp interface{}) error {
	b, err := json.Marshal(resp)
	if err != nil {
		return respondError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("failed to marshal response: %w", err),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}

func respondStatus(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func respondError(w http.ResponseWriter, statusCode int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	b := []byte(fmt.Sprintf(`{"error":"%s"}`, err))
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("failed to write error response: %w", err)
	}

	return nil
}
