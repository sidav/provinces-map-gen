package main

import (
	"province-map-generator/lib/console/tcell_console_wrapper"
	provincesgen "province-map-generator/provinces-gen"

	"github.com/gdamore/tcell/v2"
)

var (
	cw  tcell_console_wrapper.ConsoleWrapper
	gen *provincesgen.ProvincesMapGenerator
)

func main() {
	cw.Init()
	defer cw.Close()
	gen = provincesgen.NewGenerator(40, 25, 20, 2)

	gen.Init()

	draw()
	key := cw.ReadKey()

	for key != "ESCAPE" {
		if key == "ENTER" {
			gen.Generate()
		} else {
			gen.Step()
		}
		draw()
		key = cw.ReadKey()
	}
}

var provcolors = []tcell.Color{
	tcell.ColorRed,
	tcell.ColorBlue,
	tcell.ColorOrange,
	tcell.ColorGreen,
	tcell.ColorGray,
	tcell.ColorDarkMagenta,
	tcell.ColorDarkGray,
}

func draw() {
	cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkBlue)
	cw.DrawRect(0, 0, gen.Width+1, gen.Height+1)

	for x := range gen.Map {
		for y := range gen.Map[x] {
			tile := &gen.Map[x][y]
			provinceNumber := tile.ProvinceId
			switch tile.TileType {
			case provincesgen.TtypeProvince:
				cw.SetStyle(tcell.ColorBlack, provcolors[provinceNumber%len(provcolors)])
				cw.PutChar(' ', x+1, y+1)
			case provincesgen.TtypeWater:
				cw.SetStyle(tcell.ColorDarkBlue, tcell.ColorBlack)
				cw.PutChar('≈', x+1, y+1)

			}
		}
	}
	// Draw seeds:
	for id, seed := range gen.ActiveSeedsList {
		cw.SetStyle(tcell.ColorBlack, provcolors[id%len(provcolors)])
		strToPrint := string(rune('A' + id))
		cw.PutString(strToPrint, seed.X+1, seed.Y+1)
	}
	cw.FlushScreen()
}
