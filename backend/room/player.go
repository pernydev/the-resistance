package room

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Player struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	RoomID string `json:"-"`

	Sender *WSSender `json:"-"`
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

type Claims struct {
	ID     string `json:"id"`
	RoomID string `json:"room_id"`
	jwt.RegisteredClaims
}

func (p *Player) CreateAuthToken() (string, error) {
	data := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:     p.ID,
		RoomID: p.RoomID,
	})
	token, err := data.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token, err
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("token is invalid or claims could not be extracted")
	}
}
