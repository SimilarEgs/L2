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

// NewHandler func - return instance of the Handler struct
func NewHandler() *Handler {
	return &Handler{
		eventService: repository.NewEventStorage(),
	}
}

// createEvent handler - creates new event in proggram cache based on client request
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

// updateEvent handler - updates existing event in proggram cache based on client request
func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {

	// getting event request body
	event, err := h.postHelperEvent(w, r)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	// updating desired event
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

// deleteEvent handler - deletes existing event in proggram cache based on client request
func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {

	// getting event data from request body
	event, err := h.postHelperEvent(w, r)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	// deleting desired event
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

// getEventsForDay handler - retrieves existing events from the program cache based on the provided day
func (h *Handler) getEventsForDay(w http.ResponseWriter, r *http.Request) {

	// extracting query params
	event, err := h.getHelperEvent(w, r)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	// collecting data
	events, err := h.eventService.GetEvenstForDay(event.UserID, event.Date.Time.Day())
	if err != nil {
		ERROR(w, http.StatusNotFound, err)
		log.Print(err)
		return
	}

	// sending results to the client
	log.Printf("[Info] client with ID - %d retrived %d event(s)", event.UserID, len(events))
	JSON(w, http.StatusOK, resultResponse{Result: events})

}

// getEventsForWeek handler - retrieves existing events from the program cache based on the provided week (e.g - 2022-10-02: first week of the month)
func (h *Handler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {

	// extracting query params
	event, err := h.getHelperEvent(w, r)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	// collecting data
	events, err := h.eventService.GetEvenstForWeek(event.UserID, event.Date.Time)
	if err != nil {
		ERROR(w, http.StatusNotFound, err)
		log.Print(err)
		return
	}

	// sending results to the client
	log.Printf("[Info] client with ID - %d retrived %d event(s)", event.UserID, len(events))
	JSON(w, http.StatusOK, resultResponse{Result: events})

}

// getEventsForMonth handler - retrieves existing events from the program cache based on the provided month (e.g - 2022-01-02: first month of the year)
func (h *Handler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {

	// extracting query params
	event, err := h.getHelperEvent(w, r)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		log.Print(err)
		return
	}

	// collecting data
	events, err := h.eventService.GetEvenstForMonth(event.UserID, event.Date.Time.Month())
	if err != nil {
		ERROR(w, http.StatusNotFound, err)
		log.Print(err)
		return
	}

	// sending results to the client
	log.Printf("[Info] client with ID - %d retrived %d event(s)", event.UserID, len(events))
	JSON(w, http.StatusOK, resultResponse{Result: events})
}
