package game

import (
	"sync"

	"github.com/pernydev/the-resistance/backend/utils"
)

type MissionResult string

const (
	MissionResultSuccess MissionResult = "success"
	MissionResultFail    MissionResult = "fail"
)

type Mission struct {
	Index            MissionIndex    `json:"index"` // starting from 1
	ParticipantCount int             `json:"participant_count"`
	FailsRequired    int             `json:"fails_required"`
	Submissions      []MissionResult `json:"submissions"`
	Result           *MissionResult  `json:"mission_result"`

	Mutex sync.Mutex `json:"-"`

	Game *Game `json:"-"`
}

func (m *Mission) Submit(result MissionResult) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.Submissions = utils.Shuffle(append(m.Submissions, result))

	if len(m.Submissions) >= m.ParticipantCount {
		result := m.result()
		m.Result = &result
		m.Game.SetState(GameStateReveal)
	}
}

func (m *Mission) result() MissionResult {
	failsFound := 0
	for _, submission := range m.Submissions {
		if submission == MissionResultFail {
			failsFound++
		}
	}
	if failsFound > m.FailsRequired {
		return MissionResultFail
	}
	return MissionResultSuccess
}

func CreateGameMissions(game *Game) {
	missionMap := MissionMap[PlayerCount(len(game.Players))]
	for missionIndex, participantCount := range missionMap {
		failsRequired := 1
		if missionIndex == 4 && len(game.Players) >= 7 {
			failsRequired = 2
		}
		mission := Mission{
			Index:            missionIndex,
			ParticipantCount: participantCount,
			FailsRequired:    failsRequired,
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
