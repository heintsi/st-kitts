package game

import (
	"encoding/json"
	"fmt"
	"io"
)

// Game state contains all information for a single round of a game.
// After each round a new state is computed by updating it with actions
// from players.
type State struct {
	GameID  GameID
	Round	int
	Players []*Player
	// Non-exportet view in Players. Used for new state calculation.
	players map[PlayerID]*Player
}

type StateReadError struct {
	BufferLength int
	JsonLength   int
}

func (e *StateReadError) Error() string {
	return fmt.Sprintf("Buffer of lenght %v is too short for JSON of length %v",
		e.BufferLength, e.JsonLength)
}

// State's Read is used to tranform go struct to JSON.
// The JSON is written in p.
func (s *State) Read(p []byte) (n int, err error) {
	jsonBuffer, err := json.Marshal(s)
	if err != nil {
		return 0, err
	}
	n = copy(p, jsonBuffer)
	if l := len(jsonBuffer); n < l {
		err = &StateReadError{len(p), l}
	} else {
		err = io.EOF
	}
	return
}

func (s *State) allPlayersReady() bool {
	for _, player := range s.Players {
		if !player.Ready {
			return false
		}
	}
	return true
}

func (s *State) setAllPlayersNotReady() {
	for _, player := range s.Players {
		player.Ready = false
	}
}

func ExampleState(id string) *State {
	return new(State)
}
