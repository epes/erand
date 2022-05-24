package erand

import (
	"github.com/epes/goxoshiro256starstar"
	"math/rand"
)

// Random is the API for this package.
type Random struct {
	*rand.Rand
	rand.Source64
}

// New returns an instance of the erand Random API given a
// randomness Source.
func New(seed int64) *Random {
	s := goxoshiro256starstar.New(seed)

	return &Random{
		rand.New(s),
		s,
	}
}

func NewLocked(seed int64) *Random {
	s := goxoshiro256starstar.NewLocked(seed)

	return &Random{
		rand.New(s),
		s,
	}
}
