package repository

import (
	"dev11/internal/models"
	"fmt"
	"sync"
	"time"
)

// EventStorage struct - structure in which the event entity, containing users events is stored
type EventStorage struct {
	sync.RWMutex // provide concurrent access to events storage
	events       map[int]models.Event
	eventsID     int
}

// NewEventStorage func - returns a new initialized EventStorage instance
func NewEventStorage() *EventStorage {
	return &EventStorage{
		events: make(map[int]models.Event),
	}
}

// createEvent func - adds new event to the event storage
func (e *EventStorage) CreateEvent(event *models.Event) error {

	// protect for concurrent access
	e.Lock()
	defer e.Unlock()

	// check if event is already exists => if ture, return error
	if _, ok := e.events[event.EventID]; ok {
		return fmt.Errorf("[Error] event with ID - %d already exists\n", event.EventID)
	}

	// create new event
	e.events[event.EventID] = *event
	e.eventsID++

	return nil
}

// updateEvent func - updates event in the event storage
func (e *EventStorage) UpdateEvent(event *models.Event) error {

	// protect for concurrent access
	e.Lock()
	defer e.Unlock()

	// check if event is doesn't exists => if ture, return error
	if _, ok := e.events[event.EventID]; !ok {
		return fmt.Errorf("[Error] event with ID - %d doesn't exists\n", event.EventID)
	}

	// updating the event
	e.events[event.EventID] = *event

	return nil
}

// deleteEvent func - delete event in the event storage
func (e *EventStorage) DeleteEvent(eventID int) error {

	// protect for concurrent access
	e.Lock()
	defer e.Unlock()

	// check if event is doesn't exists => if ture, return error
	if _, ok := e.events[eventID]; !ok {
		return fmt.Errorf("[Error] event with ID - %d doesn't exists\n", eventID)
	}

	// delete event with provided ID
	delete(e.events, eventID)

	return nil
}

// getEvenstForDay func - returns all events in the event storage that match the given date(day)
func (e *EventStorage) GetEvenstForDay(userId int, date int) ([]models.Event, error) {

	// res - store all resualt events
	res := make([]models.Event, 0)

	// protect for concurrent read access
	e.RLock()
	defer e.RUnlock()

	// traversing through all user events -> append all matching events to res
	for _, event := range e.events {
		_, _, day := event.Date.Date()
		if event.UserID == userId && day == date {
			res = append(res, event)
		}
	}

	// checking for 0 case
	if len(res) == 0 {
		return nil, fmt.Errorf("[Info] result was not found for the entered day")
	}

	return res, nil
}

// getEvenstForDay func - returns all events in the event storage that match the given date(week)
func (e *EventStorage) GetEvenstForWeek(userId int, date time.Time) ([]models.Event, error) {

	// res - store all resualt events
	res := make([]models.Event, 0)

	inputYear, inputWeek := date.ISOWeek()

	// protect for concurrent read access
	e.RLock()
	defer e.RUnlock()

	// traversing through all user events -> append all matching events to res
	for _, event := range e.events {
		year, month := event.Date.ISOWeek()
		if event.UserID == userId && year == inputYear && month == inputWeek {
			res = append(res, event)
		}
	}

	// checking for 0 case
	if len(res) == 0 {
		return nil, fmt.Errorf("[Info] result was not found for the entered week")
	}

	return res, nil
}

// getEvenstForDay func - returns all events in the event storage that match the given date(month)
func (e *EventStorage) GetEvenstForMonth(userId int, date time.Month) ([]models.Event, error) {

	// res - store all resualt events
	res := make([]models.Event, 0)

	// protect for concurrent read access
	e.RLock()
	defer e.RUnlock()

	// traversing  through all user events -> append all matching events to res
	for _, event := range e.events {
		_, month, _ := event.Date.Date()
		if event.UserID == userId && month == date {
			res = append(res, event)
		}
	}

	// checking for 0 case
	if len(res) == 0 {
		return nil, fmt.Errorf("[Info] result was not found for the entered month")
	}

	return res, nil
}
