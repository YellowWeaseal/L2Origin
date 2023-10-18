package handler

import (
	"net/http"
	"time"
)

type eventsForDateRequest struct {
	StartDay time.Time
}

func (h *Handler) EventsForDayHandler(req *http.Request) APIResponse {
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	events, err := h.Services.GetEventsByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24))
	if err != nil {
		return h.Error(http.StatusServiceUnavailable, err)
	}

	return h.JSON(http.StatusOK, events)
}

func (h *Handler) EventsForWeekHandler(req *http.Request) APIResponse {
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	events, err := h.Services.GetEventsByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24*7))
	if err != nil {
		return h.Error(http.StatusServiceUnavailable, err)
	}

	return h.JSON(http.StatusOK, events)
}

func (h *Handler) EventsForMonthHandler(req *http.Request) APIResponse {
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	events, err := h.Services.GetEventsByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24*30))
	if err != nil {
		return h.Error(http.StatusServiceUnavailable, err)
	}

	return h.JSON(http.StatusOK, events)
}
func (data *eventsForDateRequest) parse(req *http.Request) error {
	var err error
	if err := req.ParseForm(); err != nil {
		return err
	}

	if data.StartDay, err = time.Parse("2006-01-02", req.FormValue("start_day")); err != nil {
		return err
	}

	return nil
}
