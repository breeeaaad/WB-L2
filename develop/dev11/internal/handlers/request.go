package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task11/internal/service"
	"time"
)

const (
	date  = "2006-01-02"
	month = "2006-01"
)

func parsePost(r *http.Request) (service.Event, error) {
	var event service.Event
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&event)
	if err != nil {
		return service.Event{}, err
	}
	return event, nil
}

func parseGet(r *http.Request) (time.Time, int, error) {
	date, err := time.Parse(date, r.URL.Query().Get("date"))
	if err != nil {
		return time.Time{}, 0, err
	}
	id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		return time.Time{}, 0, err
	}
	return date, id, nil
}

func parseGetMonth(r *http.Request) (time.Time, int, error) {
	date, err := time.Parse(month, r.URL.Query().Get("date"))
	if err != nil {
		return time.Time{}, 0, err
	}
	id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		return time.Time{}, 0, err
	}
	return date, id, nil
}
