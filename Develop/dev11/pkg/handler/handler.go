package handler

import (
	"dev11/pkg/repository"
	"dev11/pkg/service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	Services *service.CalendarService
	Storage  repository.CalendarStorage
	Logger   *logrus.Logger
}

func NewHandler(services *service.CalendarService, storage repository.CalendarStorage, logger *logrus.Logger) *Handler {
	return &Handler{Services: services, Storage: storage, Logger: logger}
}

type APIResponse func(resp http.ResponseWriter)

func (h *Handler) Handle(fn func(req *http.Request) APIResponse) func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Server", "GoCalendar API ")
		fn(req)(resp)
	}
}

func (h *Handler) JSON(code int, data interface{}) APIResponse {
	return h.sendJSON(code, map[string]interface{}{"result": data})
}

func (h *Handler) Error(code int, err error) APIResponse {
	return h.sendJSON(code, map[string]string{"error": err.Error()})
}

func (h *Handler) sendJSON(code int, data interface{}) APIResponse {
	var encodedData []byte
	var err error

	if data != nil {
		encodedData, err = json.Marshal(data)
		if err != nil {
			return h.Error(http.StatusInternalServerError, err)
		}
	}

	return func(resp http.ResponseWriter) {
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.Header().Set("Content-Length", strconv.Itoa(len(encodedData)))
		resp.WriteHeader(code)
		_, _ = resp.Write(encodedData)
	}
}
