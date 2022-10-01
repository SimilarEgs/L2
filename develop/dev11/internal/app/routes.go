package app

import (
	"dev11/internal/service"
	"net/http"
)

// NewRouter func - registers api endpoints
func InitRoutes(event service.EventService) {

	h := &Handler{
		eventService: event,
	}

	http.Handle("/create_event", logEvent(http.HandlerFunc(h.createEvent)))
	http.Handle("/delete_event", logEvent(http.HandlerFunc(h.updateEvent)))
	http.Handle("/update_event", logEvent(http.HandlerFunc(h.deleteEvent)))
	http.Handle("/events_for_day", logEvent(http.HandlerFunc(h.getEventsForDay)))
	http.Handle("/events_for_week", logEvent(http.HandlerFunc(h.getEventsForWeek)))
	http.Handle("/events_for_month", logEvent(http.HandlerFunc(h.getEventsForMonth)))

}
