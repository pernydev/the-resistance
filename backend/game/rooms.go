package game

import "github.com/google/uuid"

var rooms = map[string]*Room{}

func NewRoom() *Room {
	id := uuid.New().String()
	room := Room{
		ID:       id,
		Players:  map[string]*Player{},
		Settings: DefaultSettings(),
	}

	rooms[id] = &room

	return &room
}
