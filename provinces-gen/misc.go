package provincesgen

import "math/rand"

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
		if roll < wght {
			return i
		}
		roll -= wght
	}
	panic("weightedRand failed.")
}

