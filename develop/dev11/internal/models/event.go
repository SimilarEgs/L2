package models

import "time"

// Event struct - representation of event entity with all its inherent fields
type Event struct {
	EventID int  `json:"event_id"`
	UserID  int  `json:"user_id"`
	Date    Date `json:"date"`
}

// Date struct - subsidiary strucg for event fields, which store data about event date
type Date struct {
	time.Time
}

// UnmarshalJSON func - implementing json.Unmarshaler interface
func (d *Date) UnmarshalJSON(date []byte) error {

	if string(date) == "" || string(date) == `""` {
		*d = Date{time.Now()}
		return nil
	}

	t, err := time.Parse(`"`+"2006-01-02"+`"`, string(date))
	*d = Date{t}

	return err
}
