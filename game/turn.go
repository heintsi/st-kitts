package game

import (
	"encoding/json"
	"fmt"
)

// Types which are read from user input must be checkcd agaist missing
// or invalid items. The items are returned in a slice by the Check method.
type Checker interface {
	Check() []string
}

// all fields must be pointers in order to check updated fields
// when unmarshaling json.
type Turn struct {
	Player *Player
	Action *Action
}

type Action struct {
	Type string
}

func (a *Action) Check() (invalid []string) {
	if a == nil {
		invalid = append(invalid, "Action")
	} else {
		switch {
		case a.Type == "":
			invalid = append(invalid, "Action.Type")
		}
	}
	return
}

type TurnWriteError struct {
	// contains items which were not provided when submitting the turn
	MissingItems []string
}

func (e *TurnWriteError) Error() string {
	return fmt.Sprintf("JSON write failed since invalid:%v", e.MissingItems)
}

func (t *Turn) Write(p []byte) (n int, err error) {
	var ok bool
	err = json.Unmarshal(p, t)
	if err != nil {
		return 0, err
	}
	ok, err = validTurn(t)
	if ok {
		n = len(p)
	}
	return
}

// checks if all fields in Turn are ok 
// returns true and nil on success and false and error on failure
func validTurn(t *Turn) (ok bool, err error) {
	// error message will be appended with strings
	// []byte is needed since only slices can be appended.
	invalidItems := t.Check()
	if len(invalidItems) == 0 {
		ok = true
	} else {
		err = &TurnWriteError{invalidItems}
	}
	return
}

// All fields in Turn struct are checkcd. If a field is a Struct,
// it is checked also.
func (t *Turn) Check() (invalid []string) {
	if t == nil {
		invalid = append(invalid, "Turn")
	} else {
		invalid = append(invalid, t.Player.Check()...)
		invalid = append(invalid, t.Action.Check()...)
	}
	return
}
