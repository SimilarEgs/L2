package app

import (
	"dev11/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// Error representation for the client response
type errorResponse struct {
	Error string `json:"error"`
}

// Result representation for the client response
type resultResponse struct {
	Result []models.Event `json:"result"`
}

// JSON func - a wrapper for json response to client, that response statuscode and resualt
func JSON(w http.ResponseWriter, statuscode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

}

// ERROR func - a wrapper for error response to client, that response error statuscode and message of the error
func ERROR(w http.ResponseWriter, statuscode int, err error) {

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		JSON(w, statuscode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
