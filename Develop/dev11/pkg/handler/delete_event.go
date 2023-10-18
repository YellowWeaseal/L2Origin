package handler

import (
	"errors"
	"net/http"
)

type eventDeleteRequest struct {
	Id int
}

func (h *Handler) EventDeleteHandler(req *http.Request) APIResponse {
	data := &eventDeleteRequest{}
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

	if err := h.Services.RemoveEvent(event); err != nil {
		return h.Error(http.StatusServiceUnavailable, err)
	}

	return h.sendJSON(http.StatusAccepted, nil)
}

func (data *eventDeleteRequest) parse(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	eventID := req.FormValue("id")
	if eventID == "" {
		return errors.New("event ID is required")
	}

	return nil
}
