package service

import (
	"errors"
	"time"
)

func (s *Service) UpdateEvent(event Event) error {
	if _, ok := s.c[event.EventId]; ok {
		s.c[event.EventId] = event
	} else {
		return errors.New("Ивент с таким id не существует")
	}
	return nil
}

func (s *Service) CreateEvent(event Event) error {
	if _, ok := s.c[event.EventId]; !ok {
		s.c[event.EventId] = event
	} else {
		return errors.New("Ивент с таким id уже существует")
	}
	return nil
}

func (s *Service) DeleteEvent(id int) error {
	if _, ok := s.c[id]; ok {
		delete(s.c, id)
	} else {
		return errors.New("Ивент с таким id не существует")
	}
	return nil
}

func (s *Service) EventsForDay(id int, date time.Time) ([]Event, error) {
	var res []Event
	for _, v := range s.c {
		if v.Date.Format("2006-01-02") == date.Format("2006-01-02") && v.UserId == id {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("Нет ивентов за день у этого пользователя")
	}
	return res, nil
}

func (s *Service) EventsForWeek(id int, date time.Time) ([]Event, error) {
	dyear, dweek := date.ISOWeek()
	var res []Event
	for _, v := range s.c {
		year, week := v.Date.ISOWeek()
		if year == dyear && week == dweek && v.UserId == id {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("Нет ивентов за неделю у этого пользователя")
	}
	return res, nil
}

func (s *Service) EventsForMonth(id int, date time.Time) ([]Event, error) {
	var res []Event
	for _, v := range s.c {
		if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() && v.UserId == id {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("Нет ивентов за месяц у этого пользователя")
	}
	return res, nil
}
