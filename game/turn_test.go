package game

import "testing"

func TestValidTurn(t *testing.T) {
	testNilTurn(t)
	gameID, err := New()
	if err != nil {
		t.Errorf("Game creation failed when testing valid Turn")
	}
	testInvalidPlayer(t,gameID)
	testInvalidAction(t,gameID)
	End(gameID)
}

func testNilTurn(t *testing.T) {
	if ok, _ := validTurn(nil); ok {
		t.Errorf("Nil turn was considered to be valid")
	}
}

func testInvalidAction(t *testing.T, gameID GameID) {
	validPlayer := getValidPlayer()
	invalidAction := getAnInvalidAction()
	turn1 := &Turn{gameID,validPlayer,invalidAction}
	if ok, _ := validTurn(turn1); ok {
		t.Errorf("Invalid action was accepted")
	}
}

func testInvalidPlayer(t *testing.T, gameID GameID) {
	invalidPlayer := getAnInvalidPlayer()
	validAction := getValidAction()
	turn1 := &Turn{gameID,invalidPlayer,validAction}
	if ok, _ := validTurn(turn1); ok {
		t.Errorf("Invalid player was accepted")
	}
}

func getValidPlayer() PlayerID {
	return PlayerID("player")
}

func getAnInvalidPlayer() PlayerID {
	return PlayerID("")
}

func getValidAction() *Action {
	return &Action{"act"}
}

func getAnInvalidAction() *Action {
	return &Action{""}
}

func TestAction(t *testing.T) {
	var a *Action // nil
	if invalid := a.check(); len(invalid) == 0 {
		t.Errorf("Nil action was accepted")
	}
	a = &Action{""}
	if invalid := a.check(); len(invalid) == 0 {
		t.Errorf("Empty action was accepted")
	}
}

func TestWrite(t *testing.T) {
	// An existing game is required for a valid turn
	gameID, err := New("TestGame")
	if err != nil {
		t.Errorf("Creation of test game failed in Write test")
	}
	testSuccessfulWrite(t)
	testFailingWrite(t)
	End(gameID)
}

func testSuccessfulWrite(t *testing.T) {
	validTurnJson := []byte(
		`{
			"GameID" : "TestGame",
			"PlayerID" : "TestPlayer",
			"Action" : {
				"Type" : "TestAction"
			}
		}`)
	turn := new(Turn)
	_, err := turn.Write(validTurnJson)
	if err != nil {
		t.Errorf("Valid turn.Write failed with error %v", err)
	}
	if ok, _ := validTurn(turn); !ok {
		t.Errorf("Successful turn.Write created an invalid turn %v",
			turn)
	}
}

func testFailingWrite(t *testing.T) {
	// An existing game is required for a valid turn
	invalidTurnJson := []byte(
		`{
			"GameID" : "TestGame",
			"PlayerID" : "TestPlayer"
		 }`)
	turn := new(Turn)
	_, err := turn.Write(invalidTurnJson)
	if err == nil {
		t.Errorf("Invalid turn.Write succeeded with turn %v", turn)
	}
}
