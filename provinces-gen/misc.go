package provincesgen

import (
	"math/rand"
)

func percChance(perc int) bool {
	return rand.Intn(100) < perc
}

// May be extremely unefficient, depending on weightf imlementation.
func weightedRand(min, max int, weightf func(index int) int) int {
	if min == max {
		return max
	}
	totalWeights := 0
	for i := min; i <= max; i++ {
		totalWeights += weightf(i)
	}
	roll := rand.Intn(totalWeights)
	for i := min; i <= max; i++ {
		wght := weightf(i)
		if wght == 0 {
			continue
		}
		if roll < wght {
			return i
		}
		roll -= wght
	}
	panic("weightedRand failed.")
}

func weightedRandFloat(min, max int, weightf func(index int) float64) int {
	weightsSum := 0.0
	for i := min; i <= max; i++ {
		weightsSum += weightf(i)
	}
	roll := rand.Float64() * weightsSum
	for i := min; i <= max; i++ {
		w := weightf(i)
		if roll < w {
			return i
		}
		roll -= w
	}
	panic("Weighted float rand failed.")
}

