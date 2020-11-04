/*
Package ergo allows running an ergodicity economics experiment by Ole Peters
Get players to all bet $100 dollars and repeatedly flip a coin:
1. Heads their wealth increases by 50%
2. Tails their wealth decreases by 40%

A great bet on the surface.

While collective wealth increases by 5% (compounds) at every coin flip,
the vast majority of players end up skint with only few big winners; wealth
aggregation (though ultimately it's a zero sum gain). It's by no means a representation
of a the real world, but Ole Peters uses it to demonstrate the difference between _ensemble
average_ and _time average_ (and how economics does not understand that).

Import the package with
	"github.com/go-ergo/ergo"
Then play a "game," pass it the number of players, the number of coin tosses,
and the starting amount of cash of each player
	g := ergo.Play(50, 100, 1000)
You can also compute the ensemble average
	g.EnsembleAvg()
*/
package ergo
