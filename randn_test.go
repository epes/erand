package erand_test

import (
	"math/rand"
	"testing"

	"github.com/epes/erand"
)

const (
	nMax = 10
)

func TestIntn(t *testing.T) {
	r := erand.New(rand.NewSource(0))

	for i := 0; i < 1000; i++ {
		rint := r.Intn(nMax)

		if rint < 0 || rint >= nMax {
			t.FailNow()
		}
	}
}

func TestUint16n(t *testing.T) {
	r := erand.New(rand.NewSource(0))

	for i := 0; i < 1000; i++ {
		rint := r.Uint16n(nMax)

		if rint < 0 || rint >= nMax {
			t.FailNow()
		}
	}
}

func TestUint32n(t *testing.T) {
	r := erand.New(rand.NewSource(0))

	for i := 0; i < 1000; i++ {
		rint := r.Uint32n(nMax)

		if rint < 0 || rint >= nMax {
			t.FailNow()
		}
	}
}

func BenchmarkIntn(b *testing.B) {
	r := erand.New(rand.NewSource(0))

	for i := 0; i < b.N; i++ {
		r.Intn(nMax)
	}
}

func BenchmarkUint16n(b *testing.B) {
	r := erand.New(rand.NewSource(0))

	for i := 0; i < b.N; i++ {
		r.Uint16n(nMax)
	}
}

func BenchmarkUint32n(b *testing.B) {
	r := erand.New(rand.NewSource(0))

	for i := 0; i < b.N; i++ {
		r.Uint32n(nMax)
	}
}
