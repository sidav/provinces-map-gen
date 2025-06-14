package provincesgen

import "province-map-generator/lib/calculations"

type weightedCoordinate struct {
	X, Y, weight int
}

func newCoordinate(x, y int) weightedCoordinate {
	return weightedCoordinate{X: x, Y: y}
}

func newWeightedCoordinate(x, y, weight int) weightedCoordinate {
	return weightedCoordinate{X: x, Y: y, weight: weight}
}

func (c *weightedCoordinate) approxDistTo(x, y int) int {
	return calculations.GetApproxDistFromTo(c.X, c.Y, x, y)
}

func (g *ProvincesMapGenerator) SelectRandomMapCoordsByWeight(wghtFunc func(x, y int) int) *weightedCoordinate {
	var candidates []weightedCoordinate
	for x := range g.Width {
		for y := range g.Height {
			wght := wghtFunc(x, y)
			if wght > 0 {
				candidates = append(candidates, newWeightedCoordinate(x, y, wght))
			}
		}
	}
	if len(candidates) == 0 {
		return nil
	}
	index := weightedRand(0, len(candidates)-1, func(index int) int {return candidates[index].weight})
	return &candidates[index]
}

func (g *ProvincesMapGenerator) SelectRandomMapCoordsByFloatWeight(wghtFunc func(x, y int) float64) *weightedCoordinate {
	var candidates []weightedCoordinate
	var floatWeights []float64
	for x := range g.Width {
		for y := range g.Height {
			wght := wghtFunc(x, y)
			if wght > 0 {
				candidates = append(candidates, newCoordinate(x, y))
				floatWeights = append(floatWeights, wght)
			}
		}
	}
	if len(candidates) == 0 {
		return nil
	}
	index := weightedRandFloat(0, len(candidates)-1, func(index int) float64 {return floatWeights[index]})
	return &candidates[index]
}
