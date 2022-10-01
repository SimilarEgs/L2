package app

import (
	"dev11/internal/repository"
	"dev11/internal/service"
	"net/http"
)

type Handler struct {
	eventService service.EventService
}

func NewHandler() *Handler {
	return &Handler{
		eventService: repository.NewEventStorage(),
	}
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {

	JSON(w, http.StatusOK, "[OK]")
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getEventsForDay(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) eventDecodeJSON(r *http.Request) {
	
}
