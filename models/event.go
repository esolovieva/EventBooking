package models

import "time"

type Event struct {
	ID          int
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Location    string     `json:"location" binding:"required"`
	DateTime    *time.Time `json:"date" binding:"required"`
	UserID      int        `json:"user_id,omitempty"`
}

var events = []Event{}

func (e Event) Save() {
	//later add it to DB
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
