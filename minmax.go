package erand

import (
	"fmt"
	"math"
)

type minmax interface {
	EvenMinMax(int, int) int
	IncExpMinMax(int, int) int
	DecExpMinMax(int, int) int
	IncCubMinMax(int, int) int
	DecCubMinMax(int, int) int
}

// minMaxer is a function that takes a min and a max int
// and returns an int between them by random
type minMaxer = func(min int, max int) int

// EvenMinMax is a MinMaxer with even distribution
func (r *api) EvenMinMax(min int, max int) int {
	return min + r.powMinMax(min, max, 1.0)
}

// IncExpMinMax is a MinMaxer with increasing exponential distribution
// where numbers towards the minimum are less likely than numbers
// towards the maximum
func (r *api) IncExpMinMax(min int, max int) int {
	return min + r.powMinMax(min, max, 2.0)
}

// IncCubMinMax is a MinMaxer with increasing cubic distribution
// where numbers towards the minimum are less likely than numbers
// towards the maximum
func (r *api) IncCubMinMax(min int, max int) int {
	return min + r.powMinMax(min, max, 3.0)
}

// DecExpMinMax is a MinMaxer with decreasing exponential distribution
// where numbers towards the minimum are more likely than numbers
// towards the minimum
func (r *api) DecExpMinMax(min int, max int) int {
	return max - 1 - r.powMinMax(min, max, 2.0)
}

// DecCubMinMax is a MinMaxer with decreasing cubic distribution
// where numbers towards the minimum are more likely than numbers
// towards the minimum
func (r *api) DecCubMinMax(min int, max int) int {
	return max - 1 - r.powMinMax(min, max, 3.0)
}

func (r *api) powMinMax(min int, max int, pow float64) int {
	diff := max - min
	diffToPow := int(math.Pow(float64(diff), pow))
	randPowDiff := float64(r.Intn(diffToPow))
	root := int(math.Pow(randPowDiff, 1.0/pow))

	if root < 0 || root >= diff {
		panic(fmt.Sprintf("powMinMax produced out of range result | diff: %d | root: %d", diff, root))
	}

	return root
}
