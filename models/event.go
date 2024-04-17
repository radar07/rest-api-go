package models

import "time"

type Event struct {
	ID          int
	Name        string `binding: "required"`
	Description string
	Venue       string
	DateTime    time.Time
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}
