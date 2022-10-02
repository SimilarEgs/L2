package models

import "time"

// Event struct - representation of event entity with all its inherent fields
type Event struct {
	EventID int  `json:"event_id"`
	UserID  int  `json:"user_id"`
	Date    Date `json:"date"`
}

// Date struct - subsidiary struct for event fields, which store data about event date
type Date struct {
	time.Time
}
