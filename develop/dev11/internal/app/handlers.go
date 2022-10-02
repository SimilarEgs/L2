package app

import (
	"dev11/internal/models"
	"dev11/internal/repository"
	"dev11/internal/service"
	"fmt"
	"log"
	"net/http"
)

// Handler struct - injected service methods for our handlers
type Handler struct {
	eventService service.EventService
}

// Return instance of the Handler struct
func NewHandler() *Handler {
	return &Handler{
		eventService: repository.NewEventStorage(),
	}
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {

	// getting event request body
	event, err := h.postHelperEvent(w, r)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}
	// creating event from request data
	err = h.eventService.CreateEvent(event)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	// if event was successfuly created, return 201 status and resualt msg to the client
	msg := fmt.Sprintf("[Info] event with ID was %d - created", event.EventID)
	log.Println(msg)

	JSON(w, http.StatusCreated, resultResponse{Result: []models.Event{*event}})

}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {

	// getting event request body
	event, err := h.postHelperEvent(w, r)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	err = h.eventService.UpdateEvent(event)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	// if event was successfuly updated, return OK status and resualt msg to the client
	msg := fmt.Sprintf("[Info] event with ID was %d - updated", event.EventID)
	log.Print(msg)

	JSON(w, http.StatusOK, resultResponse{Result: []models.Event{*event}})

}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {

	// getting event data from request body
	event, err := h.postHelperEvent(w, r)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	err = h.eventService.DeleteEvent(event.UserID, event.EventID)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	// if event was successfuly updated, return OK status and resualt msg to the client
	msg := fmt.Sprintf("[Info] event with ID was %d - deleted", event.EventID)
	log.Print(msg)

	JSON(w, http.StatusOK, struct {
		Resualt string `json:"resualt"`
	}{msg})

}

func (h *Handler) getEventsForDay(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) eventDecodeJSON(r *http.Request) {

}
