package erand_test

import (
	"testing"

	"github.com/epes/erand"
)

const (
	mmMin = 10
	mmMax = 20
)

func BenchmarkEvenMinMax(b *testing.B) {
	r := erand.New(0)

	for i := 0; i < b.N; i++ {
		r.EvenMinMax(mmMin, mmMax)
	}
}

func BenchmarkIncExpMinMax(b *testing.B) {
	r := erand.New(0)

	for i := 0; i < b.N; i++ {
		r.IncExpMinMax(mmMin, mmMax)
	}
}

func BenchmarkDecExpMinMax(b *testing.B) {
	r := erand.New(0)

	for i := 0; i < b.N; i++ {
		r.DecExpMinMax(mmMin, mmMax)
	}
}

func BenchmarkIncCubMinMax(b *testing.B) {
	r := erand.New(0)

	for i := 0; i < b.N; i++ {
		r.IncCubMinMax(mmMin, mmMax)
	}
}

func BenchmarkDecCubMinMax(b *testing.B) {
	r := erand.New(0)

	for i := 0; i < b.N; i++ {
		r.DecCubMinMax(mmMin, mmMax)
	}
}
