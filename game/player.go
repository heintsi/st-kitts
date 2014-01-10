package game

type PlayerID string

type Player struct {
	Name  string
	Ready bool
	id    PlayerID
	turn  *Turn
}

func (id *PlayerID) check() (invalid []string) {
	if id == nil || string(*id) == "" {
		invalid = append(invalid, "PlayerID")
	}
	return
}

func (p *Player) turnReady(t *Turn) {
	p.turn = t
	p.Ready = true
}

func (p *Player) turnNotReady(t *Turn) {
	p.turn = nil
	p.Ready = false
}
