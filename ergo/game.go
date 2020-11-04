package ergo

// Play run the simulation
func Play(players, steps int, starting float32) *Game {
	game := new(Game)

	// initialise
	for i := 0; i < players; i++ {
		s := make([]Step, steps)
		s[0].Amount = starting

		for j := 0; j < len(s); j++ {
			s[j].Time = j
		}

		game.Players = append(game.Players, Player{ID: i, Steps: s})
	}

	game.RollAll(steps)

	return game
}

// EnsembleAvg avreage of all players over time
func (game *Game) EnsembleAvg() Ensemble {

	nSteps := len(game.Players[1].Steps)
	nPlayers := len(game.Players)

	steps := make([]int, nSteps)
	avg := make([]float32, nSteps)

	for s := 0; s < nSteps; s++ {
		steps[s] = s
		var total float32
		for p := 0; p < nPlayers; p++ {
			total += game.Players[p].Steps[s].Amount
		}
		avg[s] = total / float32(nPlayers)
	}

	ensemble := Ensemble{Time: steps, Amount: avg}

	return ensemble
}
