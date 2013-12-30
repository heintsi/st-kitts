package game

import "testing"

func TestPlayerCheck(t *testing.T) {
	var id *PlayerID
	out := []string{"PlayerID"}
	if invalid := id.check(); len(invalid) == 0 {
		t.Errorf("%v.check() = %v, want %v", id, invalid, out)
	}
	id = new(PlayerID)
	*id = PlayerID("")
	if invalid := id.check(); len(invalid) == 0 {
		t.Errorf("%v.check() = %v, want %v", id, invalid, out)
	}
	*id = PlayerID("aPlayer")
	out2 := []string{}
	if invalid := id.check(); len(invalid) != 0 {
		t.Errorf("%v.check() = %v, want %v", id, invalid, out2)
	}
}

func TestPlayerTurnReady(t *testing.T) {
	const pid = "p1"
	p := &Player{"Tester", false, PlayerID(pid), nil}
	turn := &Turn{GameID("test"), PlayerID(pid), &Action{"a"}}
	p.turnReady(turn)
	if p.turn != turn {
		t.Errorf("Player.turn set failed.")
	}
	if p.Ready == false {
		t.Error("Player.Ready set failed.")
	}
}

func TestPlayerTurnNotReady(t *testing.T) {
	const pid = "p1"
	turn := &Turn{GameID("test"), PlayerID(pid), &Action{"a"}}
	p := &Player{"Tester", true, PlayerID(pid), turn}
	p.turnNotReady(turn)
	if p.turn != nil {
		t.Errorf("Player.turn set failed.")
	}
	if p.Ready == true {
		t.Error("Player.Ready set failed.")
	}
}
