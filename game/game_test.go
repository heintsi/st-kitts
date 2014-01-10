package game

import "testing"

func TestNewGame(t *testing.T) {
	createGame(t)
	createNamedGame(t)
	createTwoGamesWithSameID(t)
}

func createGame(t *testing.T) {
	gameID, err := New()
	if err != nil {
		t.Errorf("game.New() failed with error: %v", err)
	} else if !gameID.exists() {
		t.Errorf("game.New() did not create game with id %v", gameID)
	} else {
		End(gameID)
	}
}

func createNamedGame(t *testing.T) {
	gameID, err := New("Some", "Name")
	if err != nil {
		t.Errorf("game.New() failed with error: %v", err)
	} else if !gameID.exists() {
		t.Errorf("game.New() did not create game with id %v", gameID)
	} else {
		End(gameID)
	}
}

func createTwoGamesWithSameID(t *testing.T) {
	gameID1, err1 := New()
	gameID2, err2 := New(gameID1.String())
	switch {
	case err1 != nil:
		t.Errorf("game.New() failed with error: %v", err1)
	case !gameID1.exists():
		t.Errorf("game.New() did not create game with id %v", gameID1)
	case err2 == nil:
		t.Errorf("game.New() did not return error while creating two "+
			"games with the same game IDs %v", gameID2)
	default:
		End(gameID1, gameID2)
	}
}

func TestEndGame(t *testing.T) {
	endOneGame(t)
	endTwoGames(t)
}

func endOneGame(t *testing.T) {
	gameID, err := New()
	if err != nil {
		t.Errorf("game.New() failed with error: %v", err)
	}
	End(gameID)
	if gameID.exists() {
		t.Errorf("game.End() did not end game %v", gameID)
	}
}

func endTwoGames(t *testing.T) {
	gameID1, err1 := New("Game1")
	gameID2, err2 := New("Game2")
	if err1 != nil || err2 != nil {
		t.Errorf("game.New() failed with errors: %v and %v", err1, err2)
	}
	End(gameID1, gameID2)
	if gameID1.exists() || gameID2.exists() {
		t.Errorf("game.End() did not end games %v and %v", gameID1, gameID2)
	}
}

func TestGameIDCheck(t *testing.T) {
	gameID, err := New()
	if err != nil {
		t.Errorf("game.New() failed with error: %v", err)
		return
	} else if invalid := gameID.check(); len(invalid) > 0 {
		t.Errorf("GameID.check() returns invalid items even though"+
			"game ID %v was just created", gameID)
	}
	End(gameID)
	if invalid := gameID.check(); len(invalid) == 0 {
		t.Errorf("GameID.check() does not return invalid items even though"+
			"game ID %v was just ended", gameID)
	}
}
