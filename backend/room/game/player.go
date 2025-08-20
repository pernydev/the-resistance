package game

type Vote bool

const (
	VoteReject  Vote = false
	VoteApprove Vote = true
)

type Side string

const (
	SideBlue Side = "blue"
	SideRed  Side = "red"
)

type Role string

const (
	RoleNormal    Role = "normal"
	RoleCommander Role = "commander"
	RoleAssasin   Role = "assasin"
	RoleBodyGuard Role = "bodyguard"
)

type RoleCard struct {
	Side Side `json:"side"`
	Role Role `json:"role"`
}

type GamePlayer struct {
	ID string `json:"id"`

	RoleCard RoleCard `json:"role_card"`

	Vote *Vote `json:"vote"`

	IsInComposition     bool `json:"is_in_composition"`
	HasCompletedMission bool `json:"has_completed_mission"`
}
