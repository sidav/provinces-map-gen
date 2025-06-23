package provincesgen

func (g *ProvincesMapGenerator) WallBetweenTwoRandomRegions() {
	id1 := weightedRand(0, len(g.Regions)-1, func(index int) int {
		if g.getRegionById(index).IsWaterRegion {
			return 0
		}
		if g.countRegionsAdjacentToRegion(index) < 3 {
			return 0
		}
		return 1
	})
	id2 := weightedRand(0, len(g.Regions)-1, func(index int) int {
		if g.getRegionById(index).IsWaterRegion || index == id1 || !g.AreRegionsAdjacent(id1, index) {
			return 0
		}
		if g.countRegionsAdjacentToRegion(index) < 2 {
			return 0
		}
		return 1
	})
	g.WallTwoAdjacentRegions(id1, id2)
}

func (g *ProvincesMapGenerator) WallTwoAdjacentRegions(id1, id2 int) {
	tilesToWall := make([]weightedCoordinate, 0)
	for x := range g.Width {
		for y := range g.Height {
			selection := percChance(1)
			if g.tileAt(x, y).belongsToProvince(id1) && g.isTileAdjacentToProvince(x, y, id2) {
				tilesToWall = append(tilesToWall, newCoordinate(x, y))

			} else if selection && g.tileAt(x, y).belongsToProvince(id2) && g.isTileAdjacentToProvince(x, y, id1) {
				tilesToWall = append(tilesToWall, newCoordinate(x, y))
			}
		}
	}
	for _, c := range tilesToWall {
		g.tileAt(c.X, c.Y).TileType = TtypeWall
	}
}
