package room

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/pernydev/the-resistance/backend/room/game"
)

func CreateGame(room *Room) {
	g := game.Game{
		ID:      room.ID,
		Players: map[string]*game.GamePlayer{},
	}
	game.CreateGameMissions(&g)
	for id := range room.Players {
		gplayer := game.GamePlayer{
			ID: id,
		}
		g.Players[id] = &gplayer
	}

	cards := getCards(room)
	fordebug, _ := json.Marshal(cards)
	fmt.Println(string(fordebug))
	cards = shuffleCards(cards)
	i := 0
	for _, p := range g.Players {
		p.RoleCard = cards[i]
		i++
	}
	room.Game = &g
	fmt.Println("updating room")
	room.Update()
	fmt.Println("updated")
}

func shuffleCards(cards []game.RoleCard) []game.RoleCard {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func getCards(room *Room) []game.RoleCard {
	roleCards := []game.RoleCard{}

	for i := 0; i < room.Settings.AmountRed; i++ {
		card := game.RoleCard{Side: game.SideRed, Role: game.RoleNormal}

		if i == 0 && room.Settings.HasAssasin {
			card.Role = game.RoleAssasin
		}

		roleCards = append(roleCards, card)
	}

	blueCount := len(room.Players) - room.Settings.AmountRed

	for i := 0; i < blueCount; i++ {
		card := game.RoleCard{Side: game.SideBlue, Role: game.RoleNormal}

		if i == 0 && room.Settings.HasCommander {
			card.Role = game.RoleCommander
		} else if i == 1 && room.Settings.HasBodyguard {
			card.Role = game.RoleBodyGuard
		}

		roleCards = append(roleCards, card)
	}

	return roleCards
}
