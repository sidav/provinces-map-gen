package perlin_noise

import (
	"math"
)

func HsvToRgb(h, s, v float64) (uint8, uint8, uint8) {
	var r, g, b, i, f, p, q, t float64
	i = math.Floor(h * 6)
	f = h*6 - i
	p = v * (1 - s)
	q = v * (1 - f*s)
	t = v * (1 - (1-f)*s)
	part := int(math.Round(i)) % 6
	switch part {
	case 0:
		r = v
		g = t
		b = p
	case 1:
		r = q
		g = v
		b = p
	case 2:
		r = p
		g = v
		b = t
	case 3:
		r = p
		g = q
		b = v
	case 4:
		r = t
		g = p
		b = v
	case 5:
		r = v
		g = p
		b = q
	}
	return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}

var baseHue = 0.0

func getSpectrumColorFor(value, min, max int) (uint8, uint8, uint8) {
	rang := max - min
	fraction := float64(value-min) / float64(rang)
	fraction *= 1.68803398875 // golden ratio
	// fmt.Printf("v %d in %d-%d is %f \n", value, min, max, fraction)
	fraction = baseHue + (fraction)
	return HsvToRgb(fraction, 0.88, 0.85)
}
