package game

type Room struct {
	ID      string             `json:"id"`
	Players map[string]*Player `json:"player"`

	HostID   string       `json:"host_id"`
	Settings GameSettings `json:"settings"`
}

func (r *Room) AddPlayer(name string) (string, error) {
	player := NewPlayer(name, r.ID)
	token, err := player.CreateAuthToken()
	r.Players[player.ID] = player
	return token, err
}
