package erand

import (
	"fmt"
	"math/rand"
	"time"
)

// RunMinMaxers runs all the MinMaxer functions through a simulation
// and outputs the results in buckets of frequency
func RunMinMaxers() {
	r := New(rand.NewSource(time.Now().UTC().UnixNano()))

	arr := [...]minMaxer{r.EvenMinMax, r.IncExpMinMax, r.DecExpMinMax, r.IncCubMinMax, r.DecCubMinMax}

	for _, fn := range arr {
		runMinmax(fn)
	}
}

func runMinmax(fn minMaxer) {

	var arr [10]int

	for i := 0; i < 10000; i++ {
		arr[fn(0, 10)]++
	}

	fmt.Println(arr)

	for i, v := range arr {
		fmt.Print(i)
		for n := 0; n < v/100; n++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
