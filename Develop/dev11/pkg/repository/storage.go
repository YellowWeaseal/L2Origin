package repository

import (
	"context"
	"dev11"
	"sync"
	"time"
)

type CalendarStorage interface {
	Save(ctx context.Context, event *dev11.Event) error
	Remove(ctx context.Context, event *dev11.Event) error
	GetEventsByPeriod(from, to time.Time) ([]*dev11.Event, error)
	GetByID(_ context.Context, id int) (*dev11.Event, error)
}

type Storage struct {
	sync.Mutex
	events map[int]*dev11.Event
}

// NewStorage Create new in-memory storage
func NewStorage() *Storage {
	eventStorage := make(map[int]*dev11.Event)
	return &Storage{events: eventStorage}
}

func (storage *Storage) Save(_ context.Context, event *dev11.Event) error {
	storage.Lock()
	defer storage.Unlock()
	storage.events[event.Id] = event

	return nil
}

func (storage *Storage) Remove(_ context.Context, event *dev11.Event) error {
	storage.Lock()
	defer storage.Unlock()
	delete(storage.events, event.Id)

	return nil
}

func (storage *Storage) GetEventsByPeriod(from, to time.Time) ([]*dev11.Event, error) {
	storage.Lock()
	defer storage.Unlock()
	result := make([]*dev11.Event, 0)

	for _, event := range storage.events {
		if event.DateFrom.Before(to) && event.DateTo.After(from) {
			result = append(result, event)
		}
	}

	return result, nil
}

func (storage *Storage) GetByID(_ context.Context, id int) (*dev11.Event, error) {
	storage.Lock()
	defer storage.Unlock()
	if event, ok := storage.events[id]; ok {
		return event, nil
	}

	return nil, nil
}
