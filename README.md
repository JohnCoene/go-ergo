# Ergodicity Economics

Reproduction of a an experiment by [Ole Peters](https://ergodicityeconomics.com/) on wealth distribution and accumulation; case against GDP.

```
go get github.com/JohnCoene/go-ergo
```

Based on an experiment in [Ergodicity Economics](https://ergodicityeconomics.com/) by [Ole Peters](https://twitter.com/ole_b_peter) which goes as follow. Get players to all bet $100 dollars and repeatedly flip a coin:

1. Heads their wealth increases by 50%
2. Tails their wealth decreases by 40%

A great bet on the surface.

> Here's a 100 dollars, flip that coin as fast as you can; I'm off to Barbados with the winnings.
> 
> -- [Rory Sutherland](https://twitter.com/rorysutherland)

While collective wealth increases by 5% (compounds) at every coin flip, the vast majority of players end up skint with only few big winners; wealth aggregation. It's by no means a representation of a the real world, but Ole Peters uses it to demonstrate the difference between _ensemble average_ and _time average_ (and how economics does not understand that).

```go
package main

import "github.com/JohnCoene/go-ergo"

func main(){
  // 50 players, 100 coin tosses
  g := ergo.Play(50, 100)

  fmt.Println(g)
}
```
