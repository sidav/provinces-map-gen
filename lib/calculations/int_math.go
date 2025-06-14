package calculations

import (
	"strconv"
)

func IntAbs(x int) int {
  y := x >> (strconv.IntSize - 1)
  return (x ^ y) - y
}

// Newton method
func IntSqrt(n int) int {
	if n == 0 {
		return 0
	}
	x := n
	for x*x > n {
		x = (x + n/x) / 2
	}
	return x
}

