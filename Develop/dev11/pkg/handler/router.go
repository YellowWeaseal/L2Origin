package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create_event", h.Handle(h.EventCreateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/update_event", h.Handle(h.EventUpdateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/delete_event", h.Handle(h.EventDeleteHandler)).Methods(http.MethodPost)
	router.HandleFunc("/events_for_day", h.Handle(h.EventsForDayHandler)).Methods(http.MethodGet)
	router.HandleFunc("/events_for_week", h.Handle(h.EventsForWeekHandler)).Methods(http.MethodGet)
	router.HandleFunc("/events_for_month", h.Handle(h.EventsForMonthHandler)).Methods(http.MethodGet)

	return router
}
