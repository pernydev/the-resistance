package game

type MissionResult string

const (
	MissionResultSuccess MissionResult = "success"
	MissionResultFail    MissionResult = "fail"
)

type Mission struct {
	Index            MissionIndex    `json:"index"` // starting from 1
	ParticipantCount int             `json:"participant_count"`
	Submissions      []MissionResult `json:"submissions"`
	Result           *MissionResult  `json:"mission_result"`

	Game *Game `json:"-"`
}

func CreateGameMissions(game *Game) {
	missionMap := MissionMap[PlayerCount(len(game.Players))]
	for missionIndex, participantCount := range missionMap {
		mission := Mission{
			Index:            missionIndex,
			ParticipantCount: participantCount,
		}
		game.Missions[missionIndex] = &mission
	}
}

// Static mess:
type PlayerCount int

var MissionMap = map[PlayerCount]map[MissionIndex]int{
	2: {
		1: 1,
		2: 1,
		3: 1,
		4: 1,
		5: 1,
	},
	5: {
		1: 2,
		2: 3,
		3: 2,
		4: 3,
		5: 3,
	},
	6: {
		1: 2,
		2: 3,
		3: 4,
		4: 3,
		5: 4,
	},
	7: {
		1: 2,
		2: 3,
		3: 3,
		4: 4,
		5: 4,
	},
	8: {
		1: 3,
		2: 4,
		3: 4,
		4: 5,
		5: 5,
	},
	9: {
		1: 3,
		2: 4,
		3: 4,
		4: 5,
		5: 5,
	},
	10: {
		1: 3,
		2: 4,
		3: 4,
		4: 5,
		5: 5,
	},
}
