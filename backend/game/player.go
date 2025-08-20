package game

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Player struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	RoomID string `json:"room_id"`
}

func NewPlayer(name, roomID string) *Player {
	id := uuid.New().String()
	player := Player{
		ID:     id,
		Name:   name,
		RoomID: roomID,
	}
	return &player
}

func (p *Player) CreateAuthToken() (string, error) {
	data := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      p.ID,
		"room_id": p.RoomID,
	})
	token, err := data.SignedString(os.Getenv("JWT_SECRET"))
	return token, err
}
