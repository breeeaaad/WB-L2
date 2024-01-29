package handlers

import (
	"encoding/json"
	"net/http"
	"task11/internal/service"
)

func ErrRes(w http.ResponseWriter, s string, status int) {
	errRes := struct {
		Err string `json:"error`
	}{Err: s}
	res, err := json.Marshal(errRes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func Success(w http.ResponseWriter, status int, event ...service.Event) {
	res, err := json.Marshal(event)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
