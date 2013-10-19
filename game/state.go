package game

import (
	"encoding/json"
	"fmt"
	"io"
)

type State struct {
	GameID uint32
	Players []*Player
}

type StateReadError struct {
	BufferLength int
	JsonLength int
}

func (e *StateReadError) Error() string {
	return fmt.Sprintf("Buffer of lenght %v is too short for JSON of length %v",
		e.BufferLength, e.JsonLength)
}

// Reading a State produces JSON encoded State in p
func (s *State) Read(p []byte) (n int, err error) {
	jsonBuffer, err := json.Marshal(s)
	if err != nil {
		return 0, err
	}
	n = copy(p,jsonBuffer)
	if l := len(jsonBuffer); n < l {
		err = &StateReadError{len(p),l}
	} else {
		err = io.EOF
	}
	return
}

func ExampleState(id uint32) *State {
	return &State{id,[]*Player{&Player{id*2}}}
}
