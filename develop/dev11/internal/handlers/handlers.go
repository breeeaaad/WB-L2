package handlers

import (
	"encoding/json"
	"net/http"
	"task11/internal/models"
)

func (h *Handler) Create_event(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(400)
		res.Write([]byte("Только метод POST"))
		return
	}
	dec := json.NewDecoder(req.Body)
	var event models.User
	if err := dec.Decode(&event); err != nil {
		res.WriteHeader(400)
		res.Write([]byte("Ошибка парсинга json"))
	}
	
}

func (h *Handler) Update_event(res http.ResponseWriter, req *http.Request) {

}

func (h *Handler) Delete_event(res http.ResponseWriter, req *http.Request) {

}

func (h *Handler) Events_for_day(res http.ResponseWriter, req *http.Request) {

}

func (h *Handler) Events_for_week(res http.ResponseWriter, req *http.Request) {

}

func (h *Handler) Events_for_month(res http.ResponseWriter, req *http.Request) {

}
