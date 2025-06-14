package provincesgen

type tileType int
const (
	TtypeEmpty tileType = iota
	TtypeProvince
	TtypeWater
	TtypeWall
)

type tile struct {
	TileType   tileType
	ProvinceId int
}

func (t *tile) belongsToProvince(id int) bool {
	return t.TileType == TtypeProvince && t.ProvinceId == id
}

// provId has no effect if tile type isn't a Province
func (t *tile) is(typ tileType, provId int) bool {
	return t.TileType == typ && (typ != TtypeProvince || t.ProvinceId == provId)
}

func (t *tile) setAsProvince(id int) {
	t.TileType = TtypeProvince
	t.ProvinceId = id
}

