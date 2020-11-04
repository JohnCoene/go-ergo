package ergo

// Play run the simulation
func Play(players, steps int) *Game {
	game := new(Game)

	// initialise
	for i := 0; i < players; i++ {
		s := make([]Step, steps)
		s[0].Amount = 100

		for j := 0; j < len(s); j++ {
			s[j].Time = j
		}

		game.Players = append(game.Players, Player{ID: i, Steps: s})
	}

	game.RollAll(steps)

	return game
}
