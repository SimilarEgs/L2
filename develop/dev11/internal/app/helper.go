package app

import (
	"dev11/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// postHelperEvent func - validate request method and parse request payload
func (h *Handler) postHelperEvent(w http.ResponseWriter, r *http.Request) (*models.Event, error) {

	// check request method and handle response error
	if r.Method != http.MethodPost {
		return nil, fmt.Errorf("[Error] invalid request method: %v", r.Method)
	}

	return h.decodeJSON(r)
}

// getHelperEvent func - validate request method and parse url params, then return parsed data
func (h *Handler) getHelperEvent(w http.ResponseWriter, r *http.Request) (*models.Event, error) {

	// check request method and handle response error
	if r.Method != http.MethodGet {
		return nil, fmt.Errorf("[Error] invalid request method: %v", r.Method)
	}

	// extracting request params
	d := r.URL.Query().Get("date")

	uID := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(uID)
	if err != nil {
		return nil, fmt.Errorf("[Error] incorrect id format: %v", r.Method)
	}

	// parse extracted date
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		return nil, fmt.Errorf("[Error] incorrect time format: %v", r.Method)
	}

	event := models.Event{
		UserID: userID,
		Date:   models.Date{Time: date},
	}

	return &event, nil
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
