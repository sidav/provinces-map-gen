package provincesgen

import (
	"math/rand"
)

func (g *ProvincesMapGenerator) doGrowthStep() bool {
	somethingChanged := false
	for provNum := range g.Regions {
		somethingChanged = g.growProvince(provNum) || somethingChanged
	}
	return somethingChanged
}

func (g *ProvincesMapGenerator) growProvince(provinceNum int) bool {
	reg := g.getRegionById(provinceNum)
	if reg.maxSize > 0 && reg.size == reg.maxSize {
		return false
	}
	growToCoords := g.SelectRandomMapCoordsByFloatWeight(
		func(x, y int) float64 {
			return g.growthCoordsWeight(x, y, provinceNum)
		})
	if growToCoords == nil {
		return false
	}
	g.growRegionInto(provinceNum, growToCoords.X, growToCoords.Y)
	return true
}

// Random weight
func (g *ProvincesMapGenerator) growthCoordsWeight(x, y, provinceNum int) float64 {
	// First of all, the coords should be empty.
	if g.Map[x][y].TileType != TtypeEmpty {
		return 0
	}
	// Second, the coords should border the province.
	if g.countSpecificAdjacentTiles(x, y, TtypeProvince, provinceNum, false) == 0 {
		return 0
	}
	// OK, now the coords should be at least applicable.
	wght := 1.0 // base weight

	// The closer to the seed we are, the better.
	distToSeed := g.getRegionById(provinceNum).SeedCoords.approxDistTo(x, y)
	wght /= float64(distToSeed * distToSeed)

	// The more adjacent tiles of same province, the better.
	adjTilesOfSameProvince := g.countSpecificAdjacentTiles(x, y, TtypeProvince, provinceNum, true)
	wght *= float64(adjTilesOfSameProvince)

	const randMaxPermutation = 0.15
	perm := rand.Float64() * 2 * randMaxPermutation
	wght = wght * (1 + perm - randMaxPermutation)

	// Water regions only: pick perimeter tiles first
	if g.getRegionById(provinceNum).IsWaterRegion {
		if x == 0 || y == 0 || x == g.Width-1 || y == g.Height - 1 {
			wght *= 1000
		}
	}
	return wght
}
