package erand

import "math/rand"

type randN interface {
	Intn(n int) int
	Uint32n(n uint32) uint32
	Uint16n(n uint16) uint16
}

// Intn returns a random int between [0, n).
func (r *api) Intn(n int) int {
	return rand.Intn(n)
}

// Uint32n returns a random uint32 between [0, n).
func (r *api) Uint32n(n uint32) uint32 {
	return uint32(rand.Intn(int(n)))
}

// Uint16n returns a random uint16 between [0, n).
func (r *api) Uint16n(n uint16) uint16 {
	return uint16(rand.Intn(int(n)))
}
