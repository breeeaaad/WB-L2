package handlers

import (
	"net/http"
	"task11/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.Create_event)
	mux.HandleFunc("/update_event", h.Update_event)
	mux.HandleFunc("/delete_event", h.Delete_event)
	mux.HandleFunc("/events_for_day", h.Events_for_day)
	mux.HandleFunc("/events_for_week", h.Events_for_week)
	mux.HandleFunc("/events_for_month", h.Events_for_month)
	middleware := Log(mux)
	return middleware
}
