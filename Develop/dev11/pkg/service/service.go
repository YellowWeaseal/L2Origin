package service

import (
	"context"
	"dev11"
	"dev11/pkg/repository"
	"time"
)

type CalendarService struct {
	storage repository.CalendarStorage
}

func NewCalendarService(storage repository.CalendarStorage) *CalendarService {
	return &CalendarService{storage: storage}
}

func (cal *CalendarService) CreateEvent(title string, from, to time.Time) (*dev11.Event, error) {
	event := dev11.NewEvent(title, from, to)
	if err := cal.storage.Save(context.Background(), event); err != nil {
		return nil, err
	}

	return event, nil
}
func (cal *CalendarService) UpdateEvent(event *dev11.Event) error {
	return cal.storage.Save(context.Background(), event)
}

// RemoveEvent Remove existing event
func (cal *CalendarService) RemoveEvent(event *dev11.Event) error {
	return cal.storage.Remove(context.Background(), event)
}

func (cal *CalendarService) GetEventsByPeriod(from, to time.Time) ([]*dev11.Event, error) {
	return cal.storage.GetEventsByPeriod(from, to)
}

func (cal *CalendarService) GetEventByID(id int) (*dev11.Event, error) {
	return cal.storage.GetByID(context.Background(), id)
}
