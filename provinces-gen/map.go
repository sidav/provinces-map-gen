package provincesgen

func (g *ProvincesMapGenerator) tileAt(x, y int) *tile {
	return &g.Map[x][y]
}

func (g *ProvincesMapGenerator) countSpecificAdjacentTiles(x, y int, typ tileType, provId int, diagonals bool) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++{
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= g.Width || y+j < 0 || y+j >= g.Height {
				continue
			}
			if i*j != 0 && !diagonals {
				continue
			}
			if g.Map[x+i][y+j].is(typ, provId) {
				count++
			}
		}
	}
	return count
}

func (g *ProvincesMapGenerator) isTileAdjacentToProvince(x, y, otherProvinceId int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++{
			if i == 0 && j == 0 || i * j != 0 {
				continue
			}
			if x+i < 0 || x+i >= g.Width || y+j < 0 || y+j >= g.Height {
				continue
			}
			if g.Map[x+i][y+j].is(TtypeProvince, otherProvinceId) {
				return true
			}
		}
	}
	return false
}

func (g *ProvincesMapGenerator) isTileAdjacentToAnotherProvince(x, y int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++{
			if i == 0 && j == 0 || i * j != 0 {
				continue
			}
			if x+i < 0 || x+i >= g.Width || y+j < 0 || y+j >= g.Height {
				continue
			}
			if g.Map[x+i][y+j].TileType == TtypeProvince && g.Map[x+i][y+j].ProvinceId != g.Map[x][y].ProvinceId {
				return true
			}
		}
	}
	return false
}

