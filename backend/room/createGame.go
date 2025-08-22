package room

import (
	"encoding/json"
	"fmt"

	"github.com/pernydev/the-resistance/backend/room/game"
	"github.com/pernydev/the-resistance/backend/utils"
)

func CreateGame(room *Room) {
	playerIDs := make([]string, 0, len(room.Players))
	for key := range room.Players {
		playerIDs = append(playerIDs, key)
	}
	g := game.Game{
		ID:          room.ID,
		Players:     map[string]*game.GamePlayer{},
		Missions:    map[game.MissionIndex]*game.Mission{},
		State:       game.GameStateRoleReveal,
		PlayerOrder: utils.Shuffle(playerIDs),
	}

	for id := range room.Players {
		gplayer := game.GamePlayer{
			ID: id,
		}
		g.Players[id] = &gplayer
	}

	cards := getCards(room)
	fordebug, _ := json.Marshal(cards)
	fmt.Println(string(fordebug))
	cards = utils.Shuffle(cards)
	i := 0
	for _, p := range g.Players {
		p.RoleCard = cards[i]
		i++
	}

	game.CreateGameMissions(&g)

	room.Game = &g
	room.Update()
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
