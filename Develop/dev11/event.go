package dev11

import (
	"sync"
	"time"
)

type Event struct {
	Id       int
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}

type UpdateEvent struct {
	Title    string    `json:"title"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

var idCounter = 0
var idMutex sync.Mutex

func NewEvent(title string, from, to time.Time) *Event {

	idMutex.Lock()
	defer idMutex.Unlock()
	event := &Event{
		Id:       idCounter,
		Title:    title,
		DateFrom: from,
		DateTo:   to,
	}
	idCounter++
	return event

}
