package handler

import (
	"errors"
	"net/http"
	"time"
)

type eventUpdateRequest struct {
	Id       int
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}

type eventUpdateResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

func (h *Handler) EventUpdateHandler(req *http.Request) APIResponse {
	data := &eventUpdateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	event, err := h.Services.GetEventByID(data.Id)
	if err != nil {
		return h.Error(http.StatusServiceUnavailable, err)
	}

	if event == nil {
		return h.Error(http.StatusNotFound, errors.New("event not found"))
	}

	event.Title = data.Title
	event.DateFrom = data.DateFrom
	event.DateTo = data.DateTo
	if err := h.Services.UpdateEvent(event); err != nil {
		return h.Error(http.StatusServiceUnavailable, err)
	}

	return h.JSON(http.StatusOK, &eventUpdateResponse{
		Id:       event.Id,
		Title:    event.Title,
		DateFrom: event.DateFrom.Format(time.RFC3339),
		DateTo:   event.DateTo.Format(time.RFC3339),
	})
}

func (data *eventUpdateRequest) parse(req *http.Request) error {
	var err error
	if err := req.ParseForm(); err != nil {
		return err
	}

	eventID := req.FormValue("id")
	if eventID == "" {
		return errors.New("event ID is required")
	}

	data.Title = req.FormValue("title")
	if data.Title == "" {
		return errors.New("event title is required")
	}

	if data.DateFrom, err = time.Parse(time.RFC3339, req.FormValue("date_from")); err != nil {
		return err
	}

	if data.DateTo, err = time.Parse(time.RFC3339, req.FormValue("date_to")); err != nil {
		return err
	}

	if data.DateFrom.After(data.DateTo) {
		return errors.New("date From can not be after Date To")
	}

	return nil
}
