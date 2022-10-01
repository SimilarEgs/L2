package service

import (
	"dev11/internal/models"
	"time"
)

// EventService interface - provide business logic for work with events
type EventService interface {
	CreateEvent(event *models.Event) error
	UpdateEvent(event *models.Event) error
	DeleteEvent(eventID int) error
	GetEvenstForDay(userId int, date int) ([]models.Event, error)
	GetEvenstForWeek(userId int, date time.Time) ([]models.Event, error)
	GetEvenstForMonth(userId int, date time.Month) ([]models.Event, error)
}
