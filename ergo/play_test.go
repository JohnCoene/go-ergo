package ergo

import (
	"testing"
)

func TestPlay(t *testing.T) {
	p := Play(100, 100, 100)
	p.EnsembleAvg()
}
