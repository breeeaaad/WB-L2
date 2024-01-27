package service

import (
	"task11/internal/models"
)

type Service map[int]models.User

func New() *Service {
	var s Service = make(map[int]models.User, 5)
	return &s
}
