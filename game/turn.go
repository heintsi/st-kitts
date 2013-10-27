package game

import (
	"encoding/json"
	"fmt"
)

// all fields must be pointers in order to check updated fields
// when unmarshaling json.
type Turn struct {
	Player *Player
	Action *Action
}

type Action struct {
	Type string
}

type TurnWriteError struct {
	// contains items which were not provided when submitting the turn
	MissingItems string
}

func (e *TurnWriteError) Error() string {
	return fmt.Sprintf("JSON write failed since missing:%s", e.MissingItems)
}

func (t *Turn) Write(p []byte) (n int, err error) {
	var ok bool
	err = json.Unmarshal(p, t)
	if err != nil {
		return 0, err
	}
	ok, err = valid(t)
	if ok {
		n = len(p)
	}
	return
}

// checks if all fields in Turn are ok i.e. not nil.
// returns true and nil on success and false and error on failure
func valid(t *Turn) (ok bool, err error) {
	// error message will be appended with strings
	// []byte is needed since only slices can be appended.
	var missingItems []byte
	switch {
		case t.Player == nil:
			missingItems = append(missingItems, " Player"...)
		case t.Action == nil:
			missingItems = append(missingItems, " Action"...)
	}
	if len(missingItems) == 0 {
		ok = true
	} else {
		err = &TurnWriteError{string(missingItems)}
	}
	return
}
