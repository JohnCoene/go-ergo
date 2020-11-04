package ergo

import (
	"log"
	"math/rand"
	"time"
)

// Step defines a time step
type Step struct {
	Time   int     `json:"stepIndex"`
	Amount float32 `json:"cash"`
}

func (s *Step) increment() {
	s.Time++
}

// Game a game object containing players
type Game struct {
	Players []Player `json:"players"`
}

// Player a player in the game and its steps
type Player struct {
	ID    int    `json:"id"`
	Steps []Step `json:"steps"`
}

// random number generator
var seed = rand.NewSource(time.Now().UnixNano())
var generator = rand.New(seed)

// coin flip
func flip() float32 {
	prop := generator.Float32()

	if prop > .5 {
		return .6
	}

	return 1.5
}

// Roll run the game one step
func (g *Game) Roll(step int) {
	if step == 0 {
		log.Fatal("Must start at step 1")
	}

	for i := 0; i < len(g.Players); i++ {
		g.Players[i].Steps[step].Amount = g.Players[i].Steps[step-1].Amount * flip()
	}
}

// RollAll run all steps
func (g *Game) RollAll(steps int) {
	for i := 1; i < steps; i++ {
		g.Roll(i)
	}
}
