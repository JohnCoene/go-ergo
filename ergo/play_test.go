package ergo

import (
	"fmt"
	"testing"
)

func TestPlay(t *testing.T) {
	p := Play(100, 100, 100)
	avg := p.EnsembleAvg()
	fmt.Println(avg)
}
