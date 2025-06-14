package perlin_noise

import "tbs4x/lib/random"

type permutationTable struct {
	table []int
}

func (pt *permutationTable) init(size int, rnd random.PRNG) {
	pt.table = make([]int, size)
	for i := 0; i < size; i++ {
		pt.table[i] = i
	}
	// now shuffle the table
	for i := size - 1; i > 0; i-- {
		j := rnd.RandInRange(0, i)
		t := pt.table[i]
		pt.table[i] = pt.table[j]
		pt.table[j] = t
	}
}

func (pt *permutationTable) getWrappedValue(index int) int {
	return pt.table[index%len(pt.table)]
}
