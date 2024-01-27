package service

import (
	"task11/internal/models"
	"time"
)

func (s *Service) Update_event(userID int, event models.Event) error {
	return nil
}

func (s *Service) Create_event(userID int, event models.Event) error {
	return nil
}

func (s *Service) Delete_event(userID int, event models.Event) error {
	return nil
}

func (s *Service) Events_for_day(id int, data time.Time) ([]models.Event, error) {
	return nil, nil
}

func (s *Service) Events_for_week(id int, data time.Time) ([]models.Event, error) {
	return nil, nil
}

func (s *Service) Events_for_month(id int, data time.Time) ([]models.Event, error) {
	return nil, nil
}
