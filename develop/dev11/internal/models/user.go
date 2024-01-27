package models

type User struct {
	Id     int           `json:"id"`
	Events map[int]Event `json:"events"`
}

func New(id int) *User {
	return &User{
		Id:     id,
		Events: make(map[int]Event, 5),
	}
}
