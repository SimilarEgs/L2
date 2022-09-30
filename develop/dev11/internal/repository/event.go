package repository

import (
	"dev11/internal/models"
	"sync"
)

// EventStorage struct - structure in which the event entity, containing users events is stored
type EventStorage struct {
	sync.RWMutex // provide concurrent access to events storage
	events       map[int]models.Event
}

// NewEventStorage func - returns a new initialized EventStorage instance
func NewEventStorage() *EventStorage {
	return &EventStorage{
		events: make(map[int]models.Event),
	}
}

