package main

import (
	"province-map-generator/lib/console/tcell_console_wrapper"
	provincesgen "province-map-generator/provinces-gen"
)

var (
	cw  tcell_console_wrapper.ConsoleWrapper
	gen *provincesgen.ProvincesMapGenerator
)

func main() {
	cw.Init()
	defer cw.Close()
	gen = provincesgen.NewGenerator(80, 35, 55, 3, 4)

	gen.Init()

	drawGeneratedMapTest()
	key := cw.ReadKey()

	for key != "ESCAPE" {
		switch key {
		case "ENTER": gen.Generate()
		case "BACKSPACE": gen.Init()
		case " ": gen.Step()
		}
		drawGeneratedMapTest()
		key = cw.ReadKey()
	}
}


