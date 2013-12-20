package game

type Player struct {
	PlayerID string
}

func (p *Player) Check() (invalid []string) {
	switch {
	case p == nil:
		invalid = append(invalid, "Player")
	case p.PlayerID == "":
		invalid = append(invalid, "Player.PlayerID")
	}
	return
}
