package game

type GameSettings struct {
	AmountEvil int `json:"amount_evil"`

	HasCommander bool `json:"has_commander"`
	HasBodyguard bool `json:"has_bodyguard"`
	HasAssasin   bool `json:"has_assasin"`
	HasSwappers  bool `json:"has_swappers"`
}

func DefaultSettings() GameSettings {
	return GameSettings{
		AmountEvil: 2,
	}
}
