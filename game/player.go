package game

type PlayerID string

type Player struct {
	PlayerID PlayerID
}

func (id *PlayerID) check() (invalid []string) {
	if id == nil || string(*id) == "" {
		invalid = append(invalid, "PlayerID")
	}
	return
}
