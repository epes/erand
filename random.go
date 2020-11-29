package erand

import "math/rand"

// Random is the API for this package.
type Random interface {
	randN
	minmax
}

type api struct {
	gorand *rand.Rand
}

// New returns an instance of the erand Random API given a
// randomness Source.
func New(src rand.Source) Random {
	return &api{
		gorand: rand.New(src),
	}
}
