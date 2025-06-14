package provincesgen

import (
	"math/rand"
	"time"
)

type ProvincesMapGenerator struct {
	Width, Height int
	Map           [][]tile

	waterThickness     int
	provincesCount int

	ActiveSeedsList []weightedCoordinate
}

func NewGenerator(width, height, provincesCount, waterThickness int) *ProvincesMapGenerator {
	return &ProvincesMapGenerator{
		Width:          width,
		Height:         height,
		provincesCount: provincesCount,
		waterThickness:     waterThickness,
	}
}

func (g *ProvincesMapGenerator) Init() {
	rand.Seed(int64(time.Now().UnixMicro()))
	// Init the map
	g.Map = make([][]tile, g.Width)
	for i := range g.Map {
		g.Map[i] = make([]tile, g.Height)
	}
	g.placeWater()
	g.placeSeeds()
}

func (g *ProvincesMapGenerator) Generate() {
	for g.doGrowthStep() {
	}
}

func (g *ProvincesMapGenerator) Step() {
	g.doGrowthStep()
}
