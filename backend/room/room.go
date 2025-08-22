package room

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pernydev/the-resistance/backend/room/game"
)

type Room struct {
	ID      string             `json:"id"`
	Players map[string]*Player `json:"players"`

	HostID   string       `json:"host_id"`
	Settings GameSettings `json:"settings"`

	Game *game.Game `json:"game"`
}

func deepCopyTree(src *Room) *Room {
	var copy Room
	jsonData, _ := json.Marshal(src)
	json.Unmarshal(jsonData, &copy)

	return &copy
}

func (r *Room) CreateGame() {
	CreateGame(r)
}

func (r *Room) Update() {
	for _, player := range r.Players {
		roomJson, err := r.MarshalFor(player.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		if player.Sender == nil {
			fmt.Println("player " + player.ID + " has no sender")
			continue
		}
		fmt.Println("player " + player.ID + " has sender")
		err = player.Sender.SendMessage(websocket.TextMessage, roomJson)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (r *Room) AddPlayer(name string) (string, error) {
	player := NewPlayer(name, r.ID)
	token, err := player.CreateAuthToken()
	r.Players[player.ID] = player
	if len(r.Players) == 1 {
		r.HostID = player.ID
	}
	r.Update()
	return token, err
}

func (r *Room) MarshalFor(playerID string) ([]byte, error) {
	copy := deepCopyTree(r)
	if copy.Game != nil {
		for id, player := range copy.Game.Players {
			if id == playerID {
				continue
			}
			player.RoleCard = game.RoleCard{Role: "redacted", Side: "redacted"}
			if r.Game.State == game.GameStateVoting {
				redactedVote := game.Vote(false)
				player.Vote = &redactedVote
			}
		}
	}

	roomJson, err := json.Marshal(copy)
	return roomJson, err
}
