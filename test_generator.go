package main

import (
	provincesgen "province-map-generator/provinces-gen"

	"github.com/gdamore/tcell/v2"
)


var provcolors = []tcell.Color{
	tcell.ColorRed,
	tcell.ColorBlue,
	tcell.ColorOrange,
	tcell.ColorRosyBrown,
	tcell.ColorGreen,
	tcell.ColorGray,
	tcell.ColorYellow,
	tcell.ColorDarkMagenta,
	tcell.ColorDarkGray,
}

func drawGeneratedMapTest() {
	cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkBlue)
	cw.DrawRect(0, 0, gen.Width+1, gen.Height+1)

	for x := range gen.Map {
		for y := range gen.Map[x] {
			tile := &gen.Map[x][y]
			cw.ResetStyle()
			chr := '?'
			provinceNumber := tile.ProvinceId
			switch tile.TileType {
			case provincesgen.TtypeProvince:
				cw.SetStyle(tcell.ColorBlack, provcolors[provinceNumber%len(provcolors)])
				chr = ' '
			case provincesgen.TtypeWater:
				cw.SetStyle(tcell.ColorDarkBlue, tcell.ColorBlack)
				chr = '≈'
			case provincesgen.TtypeWall:
				cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
				chr = '^'
			}
			cw.PutChar(chr, x+1, y+1)
		}
	}
	// Draw seeds:
	seedSym := 0
	for id, reg := range gen.Regions {
		cw.SetStyle(tcell.ColorBlack, provcolors[id%len(provcolors)])
		if reg.IsWaterRegion {
			cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkBlue)
		}
		strToPrint := string(rune('A' + seedSym))
		if !reg.IsWaterRegion {
			seedSym++
		}
		cw.PutString(strToPrint, reg.SeedCoords.X+1, reg.SeedCoords.Y+1)
	}

	cw.FlushScreen()
}

