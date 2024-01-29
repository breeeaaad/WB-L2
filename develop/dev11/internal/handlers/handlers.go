package handlers

import (
	"net/http"
	"strconv"
	"task11/internal/service"
)

func (h *Handler) Create_event(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		ErrRes(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	event, err := parsePost(req)
	if err != nil {
		ErrRes(res, err.Error(), 400)
		return
	}
	err = h.service.CreateEvent(event)
	if err != nil {
		ErrRes(res, err.Error(), 503)
		return
	}
	Success(res, http.StatusCreated, event)
}

func (h *Handler) Update_event(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		ErrRes(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	event, err := parsePost(req)
	if err != nil {
		ErrRes(res, err.Error(), 400)
		return
	}
	err = h.service.UpdateEvent(event)
	if err != nil {
		ErrRes(res, err.Error(), 503)
		return
	}
	Success(res, http.StatusOK, event)
}

func (h *Handler) Delete_event(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		ErrRes(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(req.URL.Query().Get("event_id"))
	if err != nil {
		ErrRes(res, err.Error(), 400)
		return
	}
	err = h.service.DeleteEvent(id)
	if err != nil {
		ErrRes(res, err.Error(), 503)
		return
	}
	Success(res, http.StatusNoContent, service.Event{})
}

func (h *Handler) Events_for_day(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		ErrRes(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	date, id, err := parseGet(req)
	if err != nil {
		ErrRes(res, err.Error(), 400)
		return
	}
	events, err := h.service.EventsForDay(id, date)
	if err != nil {
		ErrRes(res, err.Error(), 503)
		return
	}
	Success(res, 200, events...)
}

func (h *Handler) Events_for_week(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		ErrRes(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	date, id, err := parseGet(req)
	if err != nil {
		ErrRes(res, err.Error(), 400)
		return
	}
	events, err := h.service.EventsForWeek(id, date)
	if err != nil {
		ErrRes(res, err.Error(), 503)
		return
	}
	Success(res, 200, events...)
}

func (h *Handler) Events_for_month(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		ErrRes(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	date, id, err := parseGetMonth(req)
	if err != nil {
		ErrRes(res, err.Error(), 400)
		return
	}
	events, err := h.service.EventsForMonth(id, date)
	if err != nil {
		ErrRes(res, err.Error(), 503)
		return
	}
	Success(res, 200, events...)
}
