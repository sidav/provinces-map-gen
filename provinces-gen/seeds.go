package provincesgen

import (
	"fmt"
	"math"
	"sort"
)

const maxSeedPlacementTries = 100

func (g *ProvincesMapGenerator) placeSeeds() {
	g.placeBorderingWaterSeeds()
	maxSeedsPlaced := 0 // Needed for debug info.
	currTry := 0

	plannedSeeds := make([]weightedCoordinate, 0, g.regionsCount)
	for {
		if currTry > maxSeedPlacementTries {
			panic(fmt.Sprintf("Couldn't place %d required seeds: max of %d was placed.", g.regionsCount, maxSeedsPlaced))
		}
		currTry++
		// Picking the coords
		for {
			nextSeed := g.SelectRandomMapCoordsByWeight(func(x, y int) int {
				if g.areCoordsGoodForSeed(x, y, plannedSeeds) {
					return 1
				}
				return 0
			})
			if nextSeed != nil {
				plannedSeeds = append(plannedSeeds, *nextSeed)
				break
			} else {
			// If we can't place a seed, remove the last one, it's bad.
				currLandRegions := len(plannedSeeds)
				if currLandRegions > 0 {
					maxSeedsPlaced = max(maxSeedsPlaced, currLandRegions)
					plannedSeeds = plannedSeeds[:len(plannedSeeds)-1]
				} else {
					break
				}
			}
		}
		if len(plannedSeeds) == g.regionsCount {
			break
		}
	}

	// Sort seeds by their coordinate (this step is not neccessary).
	sort.Slice(
		plannedSeeds,
		func(i, j int) bool {
			return plannedSeeds[i].X < plannedSeeds[j].X
		},
	)

	// Apply planned seeds
	for _, p := range plannedSeeds {
		g.placeNewRegion(p.X, p.Y, 0)
	}
}

// Places seeds to grow water border on map perimeter. Ignores usual seed distances.
func (g *ProvincesMapGenerator) placeBorderingWaterSeeds() {
	const meanMargin = 4
	for x := range g.Width {
		for y := range g.Height {
			if (x == 0 || y == 0 || x == g.Width-1 || y == g.Height-1) && (x + y) % meanMargin == 0 {
				g.placeNewRegion(x, y, (500*meanMargin+50)/100)
				g.getLastRegion().IsWaterRegion = true
			}
		}
	}
}

var minDistBetweenSeeds = -1

func (g *ProvincesMapGenerator) areCoordsGoodForSeed(x, y int, plannedSeeds []weightedCoordinate) bool {
	borderOffset := min(g.Width, g.Height)/5
	if x < borderOffset || x >= g.Width - borderOffset || y < borderOffset || y >= g.Height - borderOffset {
		return false
	}
	if g.Map[x][y].TileType != TtypeEmpty {
		return false
	}
	if minDistBetweenSeeds == -1 {
		provinceRadius := math.Sqrt(float64((g.Width-2)*(g.Height-2)) / (math.Pi * float64(g.regionsCount)))
		minDistBetweenSeeds = int(1.5 * provinceRadius)
	}
	// Check against existing seeds
	const minDistToWater = 3
	for _, as := range g.Regions {
		if as.IsWaterRegion && as.SeedCoords.approxDistTo(x, y) < minDistToWater {
			return false
		} 
	}
	// Check against planned seeds
	for _, pseed := range plannedSeeds {
		if pseed.approxDistTo(x, y) < minDistBetweenSeeds {
			return false
		}
	}

	return true
}
