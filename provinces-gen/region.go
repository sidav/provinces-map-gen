package provincesgen

type region struct {
	Id            int
	SeedCoords    weightedCoordinate
	IsWaterRegion bool

	// Internal
	size    int
	maxSize int
}

func (g *ProvincesMapGenerator) placeNewRegion(seedX, seedY, maxSize int) {
	newId := len(g.Regions)
	g.Regions = append(g.Regions, &region{
		Id:         newId,
		SeedCoords: weightedCoordinate{X: seedX, Y: seedY},
		maxSize:    maxSize,
	})
	g.growRegionInto(newId, seedX, seedY)
}

func (g *ProvincesMapGenerator) getRegionById(id int) *region {
	reg := g.Regions[id]
	if reg.Id != id {
		panic("Region IDs inconsistency detected.")
	}
	return reg
}

func (g *ProvincesMapGenerator) growRegionInto(regionId, x, y int) {
	g.Regions[regionId].size++
	g.Map[x][y].setAsProvince(regionId)
}

func (g *ProvincesMapGenerator) doesRegionBorderMap(regionId int) bool {
	for x := range g.Width {
		if g.tileAt(x, 0).belongsToProvince(regionId) || g.tileAt(x, g.Height-1).belongsToProvince(regionId) {
			return true
		}
	}
	for y := range g.Height {
		if g.tileAt(0, y).belongsToProvince(regionId) || g.tileAt(g.Width-1, y).belongsToProvince(regionId) {
			return true
		}
	}
	return false
}

func (g *ProvincesMapGenerator) fillRegionWithWater(regionId int) {
	for x := range g.Width {
		for y := range g.Height {
			if g.Map[x][y].belongsToProvince(regionId) {
				g.Map[x][y].TileType = TtypeWater
			}
		}
	}
	g.getRegionById(regionId).IsWaterRegion = true
}

func (g *ProvincesMapGenerator) getLastRegion() *region {
	return g.Regions[len(g.Regions)-1]
}

func (g *ProvincesMapGenerator) removeLastRegion() {
	for x := range g.Width {
		for y := range g.Height {
			if g.Map[x][y].ProvinceId == g.getLastRegion().Id {
				g.Map[x][y].TileType = TtypeEmpty
			}
		}
	}
	g.Regions = g.Regions[:len(g.Regions)-1]
}

func (g *ProvincesMapGenerator) isRegionAdjacentToWater(thisRegionId int) bool {
	for x := range g.Width {
		for y := range g.Height {
			if g.tileAt(x, y).TileType == TtypeWater && g.isTileAdjacentToProvince(x, y, thisRegionId) {
				return true
			}
		}
	}
	return false
}

func (g *ProvincesMapGenerator) AreRegionsAdjacent(thisRegionId int, otherRegionId int) bool {
	for x := range g.Width {
		for y := range g.Height {
			if g.tileAt(x, y).belongsToProvince(thisRegionId) && g.isTileAdjacentToProvince(x, y, otherRegionId) {
				return true
			}
		}
	}
	return false
}

func (g *ProvincesMapGenerator) countRegionsAdjacentToRegion(regId int) int {
	otherRegions := make(map[int] bool, 0)
	for x := range g.Width {
		for y := range g.Height {
			if g.tileAt(x, y).TileType != TtypeProvince {
				continue
			}
			if !g.tileAt(x, y).belongsToProvince(regId) && g.isTileAdjacentToProvince(x, y, regId) {
				if !otherRegions[g.tileAt(x, y).ProvinceId] {
					otherRegions[g.tileAt(x, y).ProvinceId] = true
				}
			}
		}
	}
	return len(otherRegions)
}

func (g *ProvincesMapGenerator) countPlacedRegions() int {
	return len(g.Regions)
}

func (g *ProvincesMapGenerator) countPlacedLandRegions() int {
	count := 0
	for _, r := range g.Regions {
		if !r.IsWaterRegion {
			count++
		}
	}
	return count
}
