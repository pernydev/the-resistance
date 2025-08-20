package room

type GameSettings struct {
	AmountRed int `json:"amount_red"`

	HasCommander bool `json:"has_commander"`
	HasBodyguard bool `json:"has_bodyguard"`
	HasAssasin   bool `json:"has_assasin"`
	HasSwappers  bool `json:"has_swappers"`
}

func DefaultSettings() GameSettings {
	return GameSettings{
		AmountRed: 1,
	}
}
