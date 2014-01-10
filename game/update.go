package game

func stateGenerator(id GameID) {
	var t *Turn
	for {
		t = <-TurnChannel
		go updateState(t)
	}
}

func updateState(t *Turn) {
	games.mutex.Lock()
	defer games.mutex.Unlock()
	state := games.inPlay[t.GameID]
	state.players[t.PlayerID].turnReady(t)
	if state.allPlayersReady() {
		computeNew(state)
	}
}

func computeNew(s *State) {
	s.setAllPlayersNotReady()
	s.Round++
}
