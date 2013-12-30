package game

import (
	"encoding/json"
	"fmt"
)

// Types which are read from user input must be checked against missing
// or invalid items. The items are returned in a slice by the check method.
type checker interface {
	check() []string
}

// Turn structure contains fields which will be parsed from JSON
// and eventually submitted to game state generator.
type Turn struct {
	GameID   GameID
	PlayerID PlayerID
	// Pointer value *Action needed for json unmarshaling.
	Action *Action
}

// Action contains player's move.
type Action struct {
	Type string
}

func (a *Action) check() (invalid []string) {
	switch {
	case a == nil:
		invalid = append(invalid, "Action")
	case a.Type == "":
		invalid = append(invalid, "Action.Type")
	}
	return
}

// TurnWriteError is used to indicate missing items in an invalid turn.
type TurnWriteError struct {
	MissingItems []string
}

func (e *TurnWriteError) Error() string {
	return fmt.Sprintf("JSON write failed since invalid:%v", e.MissingItems)
}

// Tranform JSON in p to a turn structure. All required fields in turn must be
// provided. If some fields are missing an error containing the missing fields
// is returned.
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

// Checks if a turn is valid i.e. all required information was provided.
// If turn was valid returns (true, nil) and if not (false, error).
func validTurn(t *Turn) (ok bool, err error) {
	invalidItems := t.check()
	if len(invalidItems) == 0 {
		ok = true
	} else {
		err = &TurnWriteError{invalidItems}
	}
	return
}

func (t *Turn) check() (invalid []string) {
	// Turn is Checked by checking its components
	if t == nil {
		invalid = append(invalid, "Turn")
		return
	}
	invalid = append(invalid, t.GameID.check()...)
	invalid = append(invalid, t.PlayerID.check()...)
	invalid = append(invalid, t.Action.check()...)
	return
}

// Sends turn struct to procedure which computes a new game state.
func (t *Turn) Submit() {
	// the turn should be ok if checked properly. If GameID is somehow
	// wrong, there is not much to be done here.
	if !t.GameID.exists() {
		panic(fmt.Sprintf("Invalid GameID %v", t.GameID))
	}
	TurnChannel <- t
}
