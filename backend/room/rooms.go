package room

import (
	"fmt"

	"github.com/google/uuid"
)

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

func GetRoom(id string) (*Room, error) {
	room, ok := rooms[id]
	if !ok {
		return nil, fmt.Errorf("no room found")
	}

	return room, nil
}
