package game

type MissionIndex uint8

const (
	MissionIndexRoleReveal MissionIndex = iota // Not in a mission
	MissionIndex1
	MissionIndex2
	MissionIndex3
	MissionIndex4
	MissionIndex5
)

type GameState string

const (
	GameStateRoleReveal          GameState = "role_reveal"
	GameStateCompositionCreation GameState = "composition_creation"
	GameStateVoting              GameState = "voting"
	GameStateVotingReveal        GameState = "voting_reveal"
	GameStateMission             GameState = "mission"
	GameStateReveal              GameState = "reveal"
)

type Game struct {
	ID      string       `json:"-"`
	Mission MissionIndex `json:"round"`
	State   GameState    `json:"state"`

	Missions map[MissionIndex]*Mission `json:"missions"`
	Players  map[string]*GamePlayer    `json:"players"`

	CompositionCreator int `json:"composition_creator"`
}

func (g *Game) SetState(state GameState) {
	g.State = state
}
