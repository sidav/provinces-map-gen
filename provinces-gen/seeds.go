package provincesgen

import (
	"fmt"
	"math"
	"math/rand"
	"province-map-generator/lib/calculations"
	"sort"
)

const maxSeedPlacementTries = 100000
const maxSingleSeedPlacementTries = 1000

func (g *ProvincesMapGenerator) placeSeeds() {
	maxSeedsPlaced := 0 // Needed for debug info.
	currTry := 0
	for {
		if currTry > maxSeedPlacementTries {
			panic(fmt.Sprintf("Couldn't place %d required seeds: max of %d was placed.", g.provincesCount, maxSeedsPlaced))
		}
		currTry++
		currSeedTry := 0
		// Picking the coords
		for {
			currSeedTry++

			x, y := g.randomSeedCoords()
			if g.areCoordsGoodForSeed(x, y) {
				g.ActiveSeedsList = append(g.ActiveSeedsList, weightedCoordinate{x, y, -1})
				break
			}
			// If we can't place a seed, remove the last one, it's bad.
			if currSeedTry >= maxSingleSeedPlacementTries {
				if len(g.ActiveSeedsList) > 0 {
					maxSeedsPlaced = max(maxSeedsPlaced, len(g.ActiveSeedsList))
					g.ActiveSeedsList = g.ActiveSeedsList[:len(g.ActiveSeedsList)-1]
				} else {
					panic("Unable to find good first seed.")
				}
			}
		}
		// Sort seeds by their coordinate (this step is not neccessary).
		sort.Slice(g.ActiveSeedsList, func(i, j int) bool { return g.ActiveSeedsList[i].X < g.ActiveSeedsList[j].X })

		// Place the seeds on the map.
		if len(g.ActiveSeedsList) == g.provincesCount {
			break
		}
	}
	for i, s := range g.ActiveSeedsList {
		g.Map[s.X][s.Y].TileType = TtypeProvince
		g.Map[s.X][s.Y].ProvinceId = i
	}
}

func (g *ProvincesMapGenerator) randomSeedCoords() (int, int) {
	const borderOffset = 2
	return borderOffset + rand.Intn(g.Width-2*borderOffset), borderOffset + rand.Intn(g.Height-2*borderOffset)

}

var minDistBetweenSeeds = -1

func (g *ProvincesMapGenerator) areCoordsGoodForSeed(x, y int) bool {
	if g.Map[x][y].TileType != TtypeEmpty {
		return false
	}
	if minDistBetweenSeeds == -1 {
		provinceRadius := math.Sqrt(float64(g.Width*g.Height)/(math.Pi*float64(g.provincesCount)))
		minDistBetweenSeeds = int(1.5 * provinceRadius)
	}
	for _, as := range g.ActiveSeedsList {
		if calculations.AreCoordsInRange(as.X, as.Y, x, y, minDistBetweenSeeds) {
			return false
		}
	}
	return true
}

