package service

import "time"

type Event struct {
	UserId      int       `json:"user_id"`
	EventId     int       `json:"event_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Service struct {
	c map[int]Event
}

func New() *Service {
	return &Service{
		c: make(map[int]Event, 10),
	}
}
