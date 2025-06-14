package provincesgen

import (
	"math/rand"
	"province-map-generator/lib/calculations"
)

func (g *ProvincesMapGenerator) doGrowthStep() bool {
	somethingChanged := false
	for provNum := range g.provincesCount {
		somethingChanged = g.growProvince(provNum) || somethingChanged
	}
	return somethingChanged
}

func (g *ProvincesMapGenerator) growProvince(provinceNum int) bool {
	// First, create a list of applicable coords to grow into
	// Third coordinate is a weight
	growthCandidatesList := make([]weightedCoordinate, 0)
	for x := range g.Width {
		for y := range g.Height {
			wght := g.growthCoordsWeight(x, y, provinceNum)
			if wght != 0 {
				growthCandidatesList = append(growthCandidatesList, weightedCoordinate{x, y, wght})
			}
		}
	}
	if len(growthCandidatesList) == 0 {
		return false
	}
	selectedCoordsIndex := weightedRand(0, len(growthCandidatesList)-1,
		func(i int) int {
			return growthCandidatesList[i].weight
		})
	selectedCoords := growthCandidatesList[selectedCoordsIndex]
	g.Map[selectedCoords.X][selectedCoords.Y].setAsProvince(provinceNum)
	return true
}

// Random weight
func (g *ProvincesMapGenerator) growthCoordsWeight(x, y, provinceNum int) int {
	// First of all, the coords should be empty.
	if g.Map[x][y].TileType != TtypeEmpty {
		return 0
	}
	// Second, the coords should border the province.
	if g.countSpecificAdjacentTiles(x, y, TtypeProvince, provinceNum, false) == 0 {
		return 0
	}
	// OK, now the coords should be at least applicable.
	wght := calculations.GetApproxDistFromTo(g.Width, g.Height, 0, 0) // base weight - max possible distance
	// Random permutation of base weight.
	const randomPermPercent = 35
	perm := 100 - randomPermPercent + rand.Intn(2*randomPermPercent+1)
	wght = (wght*perm + 50) / 100
	// The closer to the seed we are, the better.
	distToSeed := g.ActiveSeedsList[provinceNum].approxDistTo(x, y)
	wght /= distToSeed
	// The more adjacent tiles of same province, the better.
	adjTilesOfSameProvince := g.countSpecificAdjacentTiles(x, y, TtypeProvince, provinceNum, true)
	wght *= adjTilesOfSameProvince
	return max(wght, 1)
}
