// Package game contains game logic and front-end JSON transformations.
//
// Main idea is players to submit their moves in JSON which is parsed
// and collected. When all players have submitted, a new state of the game
// is computed. The state is then tranformed back in JSON and provided to
// players.
package game

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	// Game ID is used to identify each game instance.
	GameID string
	// Invalid game id error is used when trying to get a state not in play
	// or creating a game with game id currently in use.
	InvalidGameIDError GameID
)

var (
	// Games currently played are stored in games.InPlay which is made
	// suitable for concurrent use with reader/writer mutex.
	games = struct {
		inPlay map[GameID]*State
		mutex  sync.RWMutex
	}{inPlay: make(map[GameID]*State)}
	// Turn channel is used to input players' turns
	TurnChannel chan *Turn
)

// A new game is created and a unique game id is assigned if successful. 
// The game id can be input as arguments or one can let the procedure
// generate an ID by calling New with no arguments.
// Function accepts zero or many GameIDs as input. With multiple inputs,
// concatenation of the arguments is the game id used. If requested game id
// is in use, undefined id and an InvalidGameIDError are returned.
func New(GameIDs ...string) (id GameID, err error) {
	if len(GameIDs) == 0 {
		// FIXME: generation, this will suffice for now even though
		// two same ids are generated if two games are created within
		// one nanosecond.
		id = gameIDFromString(fmt.Sprintf("g%s", time.Now().UnixNano()))
	} else {
		id = gameIDFromString(strings.Join(GameIDs, ""))
	}
	if id.exists() {
		err = InvalidGameIDError(id)
	} else {
		initializeGameState(id)
	}
	return
}

func initializeGameState(id GameID) {
	state := new(State)
	state.GameID = id
	// launch turn updater here?
	games.mutex.Lock()
	defer games.mutex.Unlock()
	games.inPlay[id] = state
}

func (id *GameID) String() string {
	return string(*id)
}

func (id *GameID) exists() (ok bool) {
	games.mutex.RLock()
	defer games.mutex.RUnlock()
	_, ok = games.inPlay[*id]
	return
}

func gameIDFromString(strID string) GameID {
	return GameID(strID)
}

func (e InvalidGameIDError) Error() string {
	return fmt.Sprintf("Invalid game id %s", e)
}

// Retreives a copy of game state corresponding to game id. If not found
// returns an InvalidGameIDError.
func (id *GameID) GetState() (s *State, err error) {
	games.mutex.RLock()
	defer games.mutex.RUnlock()
	state, ok := games.inPlay[*id]
	if !ok {
		err = InvalidGameIDError(*id)
	} else {
		s = new(State)
		*s = *state
	}
	return
}

// Checks if a game id is valid, i.e. is found in games map.
func (id *GameID) check() (invalid []string) {
	if !id.exists() {
		invalid = append(invalid, "GameID")
	}
	return
}
