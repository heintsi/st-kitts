package game

type Player struct {
	PlayerID string
}

func (p *Player) Check() (invalid []string) {
	if p == nil {
		invalid = append(invalid, "Player")
	} else {
		switch {
		case p.PlayerID == "":
			invalid = append(invalid, "Player.PlayerID")
		}
	}
	return
}
