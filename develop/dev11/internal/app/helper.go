package app

import (
	"dev11/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// getHelperEvent func - validate request method and parse request payload
func (h *Handler) postHelperEvent(w http.ResponseWriter, r *http.Request) (*models.Event, error) {

	// check request method and handle response error
	if r.Method != http.MethodPost {
		return nil, fmt.Errorf("[Error] invalid request method: %v", r.Method)
	}

	return h.decodeJSON(r)
}

// getHelperEvent func - validate request method and parse request payload
func (h *Handler) getHelperEvent(w http.ResponseWriter, r *http.Request) (*models.Event, error) {

	// check request method and handle response error
	if r.Method != http.MethodGet {
		return nil, fmt.Errorf("[Error] invalid request method: %v", r.Method)
	}

	return h.decodeJSON(r)
}

// decodeJSON func - decode request body into event model and return result
func (h *Handler) decodeJSON(r *http.Request) (*models.Event, error) {

	event := models.Event{}

	// decode request payload
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
