package perlin_noise

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func CreatePerlinNoiseImage(noiseMap [][]float64) {
	img := image.NewRGBA(image.Rect(0, 0, len(noiseMap), len(noiseMap[0])))

	for x := 0; x < len(noiseMap); x++ {
		for y := 0; y < len(noiseMap[x]); y++ {
			col := color.RGBA{A: 255}
			fraction := noiseMap[x][y]
			// fmt.Printf("Fraction %.2f, res %d\n", fraction, scaleToRange(fraction, 200, 255))
			switch {
			//case fraction < 0.05:
			//	break
			case fraction < 0.5:
				col.B = scaleToRange(fraction, 16, 255)
				break
			case fraction < 0.8:
				col.G = scaleToRange(fraction, 16, 200)
				break
			case fraction < 0.888888:
				col.R = scaleToRange(fraction, 16, 96)
				col.G = scaleToRange(fraction, 16, 96)
				break
			default:
				col.R = scaleToRange(fraction, 210, 255)
				col.G = scaleToRange(fraction, 200, 255)
				col.B = scaleToRange(fraction, 190, 255)
			}
			//col.R = scaleToRange(fraction, 0, 255)
			//col.G = scaleToRange(fraction, 0, 255)
			//col.B = scaleToRange(fraction, 0, 255)

			//if fraction > 0 {
			//	col.R, col.G, col.B = HsvToRgb(fraction/4, 0.88, 0.85)
			//}

			img.Set(x, y, col)
		}
	}
	f, err := os.Create("perlin.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func scaleToRange(v float64, min, max uint8) uint8 {
	return uint8(math.Round(float64(max-min)*v)) + min
}
