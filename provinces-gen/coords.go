package provincesgen

import "province-map-generator/lib/calculations"

type weightedCoordinate struct {
	X, Y, weight int
}

func (c *weightedCoordinate) approxDistTo(x, y int) int {
	return calculations.GetApproxDistFromTo(c.X, c.Y, x, y)
}

