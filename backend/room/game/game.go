package game

import "sync"

type MissionIndex uint8

const (
	MissionIndexRoleReveal MissionIndex = iota // Not in a mission
	MissionIndex1
	MissionIndex2
	MissionIndex3
	MissionIndex4
	MissionIndex5
	MissionAssasinGuess // Blue won, assasin has to guess commander
	MissionEnd          // End of game
)

type GameState string

const (
	GameStateRoleReveal          GameState = "role_reveal"
	GameStateCompositionCreation GameState = "composition_creation"
	GameStateVoting              GameState = "voting"
	GameStateVotingReveal        GameState = "voting_reveal"
	GameStateMission             GameState = "mission"
	GameStateReveal              GameState = "reveal"
	GameStateAssasinGuess        GameState = "assasin_guess"
	GameStateEnd                 GameState = "end"
)

type Game struct {
	ID      string       `json:"-"`
	Mission MissionIndex `json:"round"`
	State   GameState    `json:"state"`

	Missions map[MissionIndex]*Mission `json:"missions"`
	Players  map[string]*GamePlayer    `json:"players"`

	PlayerOrder                []string `json:"player_order"`
	CurrentPlayer              int      `json:"current_player"`
	RejectedCompositions       int      `json:"rejected_compositions"`
	CurrentCompositionRejected bool     `json:"current_composition_rejected"`

	Mutex sync.Mutex `json:"-"`
}

func (g *Game) SetState(state GameState) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	g.State = state
	if state == GameStateVoting && g.CurrentCompositionRejected {
		g.CurrentCompositionRejected = false
	}
}

func (g *Game) AmountInComposition() int {
	amount := 0
	for _, player := range g.Players {
		if player.IsInComposition == true {
			amount++
		}
	}
	return amount
}

func (g *Game) NextMission() {
	g.Mission++
	g.SetState(GameStateCompositionCreation)
	for _, player := range g.Players {
		player.IsInComposition = false
		player.Vote = nil
	}
}

func (g *Game) NextPlayer() {
	g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.PlayerOrder)
}

func (g *Game) hasAllVoted() bool {
	for _, player := range g.Players {
		if player.Vote == nil {
			return false
		}
	}
	return true
}

func (g *Game) Vote(playerID string, vote Vote) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	g.Players[playerID].Vote = &vote
	if g.hasAllVoted() {
		g.SetState(GameStateVotingReveal)
	}
}
