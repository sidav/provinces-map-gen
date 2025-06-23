package provincesgen

import (
	"math/rand"
	"time"
)

type ProvincesMapGenerator struct {
	Width, Height int
	Map           [][]tile

	regionsCount      int
	waterRegionsCount int
	wallsCount        int

	Regions []*region
}

func NewGenerator(width, height, regionsCount, waterRegionsCount, wallsCount int) *ProvincesMapGenerator {
	return &ProvincesMapGenerator{
		Width:             width,
		Height:            height,
		regionsCount:      regionsCount,
		waterRegionsCount: waterRegionsCount,
		wallsCount:        wallsCount,
	}
}

// Also can be used to reset
func (g *ProvincesMapGenerator) Init() {
	rand.Seed(int64(time.Now().UnixMicro()))
	// Init the map
	g.Map = make([][]tile, g.Width)
	for i := range g.Map {
		g.Map[i] = make([]tile, g.Height)
	}
	g.Regions = nil
	g.placeSeeds()
	g.makeVoronoiRegions()
	for i := range g.Regions {
		if g.doesRegionBorderMap(i) {
			g.fillRegionWithWater(i)
		}
	}
}

func (g *ProvincesMapGenerator) Generate() {
	for g.doGrowthStep() {
	}
	for id, r := range g.Regions {
		if r.IsWaterRegion {
			g.fillRegionWithWater(id)
		}
	}
	for range g.waterRegionsCount {
		for {
			idToFill := rand.Intn(g.countPlacedRegions())
			if g.getRegionById(idToFill).IsWaterRegion {
				continue
			}
			if g.countRegionsAdjacentToRegion(idToFill) == 2 {
				continue
			}
			if g.isRegionAdjacentToWater(idToFill) && percChance(50) {
				continue
			}
			g.fillRegionWithWater(idToFill)
			break
		}
	}
	for range g.wallsCount {
		g.WallBetweenTwoRandomRegions()
	}
}

func (g *ProvincesMapGenerator) Step() {
	g.doGrowthStep()
}
