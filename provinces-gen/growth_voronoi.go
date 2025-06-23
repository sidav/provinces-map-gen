package provincesgen

func (g *ProvincesMapGenerator) makeVoronoiRegions() {
	for x := range g.Width {
		for y := range g.Height {
			g.tileAt(x, y).ProvinceId = g.regionWithClosestSeedToCoords(x, y).Id
			g.tileAt(x, y).TileType = TtypeProvince
		}
	}
}

func (g *ProvincesMapGenerator) regionWithClosestSeedToCoords(x, y int) *region {
	var closest *region = nil
	for _, r := range g.Regions {
		if closest == nil || closest.SeedCoords.sqDistTo(x, y) > r.SeedCoords.sqDistTo(x, y) {
			closest = r
		}
	}
	return closest
}

