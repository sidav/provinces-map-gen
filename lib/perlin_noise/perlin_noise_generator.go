package perlin_noise

import (
	"math"
	"tbs4x/lib/random"
)

var permTable permutationTable

func getConstantVector(v int) *vector {
	switch v & 3 {
	case 0:
		return &vector{1, 1}
	case 1:
		return &vector{-1, 1}
	case 2:
		return &vector{-1, -1}
	case 3:
		return &vector{1, -1}
	}
	panic("Wat")
}

func fade(t float64) float64 {
	return ((6*t-15)*t + 10) * t * t * t
}

func linearInterpolation(t, a1, a2 float64) float64 {
	return a1 + t*(a2-a1)
}

func hashCorner(x, y int) int {
	return permTable.getWrappedValue(permTable.getWrappedValue(x) + y)
}

func noise2d(x, y float64) float64 {
	ix := int(x) % len(permTable.table)
	iy := int(y) % len(permTable.table)

	xf := x - math.Floor(x)
	yf := y - math.Floor(y)

	topRight := &vector{xf - 1, yf - 1}
	topLeft := &vector{xf, yf - 1}
	botRight := &vector{xf - 1, yf}
	botLeft := &vector{xf, yf}

	valTopRight := hashCorner(ix+1, iy+1)
	valTopLeft := hashCorner(ix, iy+1)
	valBotRight := hashCorner(ix+1, iy)
	valBotLeft := hashCorner(ix, iy)

	// fmt.Printf("Coords (%d, %d): xf, yf are (%.1f, %.1f); %d, %d, %d, %d", ix, iy, xf, yf, valTopRight, valTopLeft, valBotRight, valBotLeft)

	dotTR := topRight.dotProduct(getConstantVector(valTopRight))
	dotTL := topLeft.dotProduct(getConstantVector(valTopLeft))
	dotBR := botRight.dotProduct(getConstantVector(valBotRight))
	dotBL := botLeft.dotProduct(getConstantVector(valBotLeft))

	// fmt.Printf("  DPs: %.1f, %.1f, %.1f, %.1f \n", dotTR, dotTL, dotBR, dotBL)

	u := fade(xf)
	v := fade(yf)
	return linearInterpolation(u, linearInterpolation(v, dotBL, dotTL), linearInterpolation(v, dotBR, dotTR))
}

func GeneratePerlinNoiseMap(w, h, noiseTileSize, FBMIterations int, rnd random.PRNG) [][]float64 {
	permTableSize := 10 // int(float64(w*h) * scale * scale)
	permTable.init(permTableSize, rnd)
	scale := 1 / float64(noiseTileSize)
	noiseMap := make([][]float64, w)
	for x := 0; x < w; x++ {
		noiseMap[x] = make([]float64, h)
		for y := 0; y < h; y++ {
			noise := 0.0
			if FBMIterations > 0 { // fractal brownian motion
				currScale := scale
				a := 1.0
				for o := 0; o < FBMIterations; o++ {
					noise += a * noise2d(float64(x)*currScale, float64(y)*currScale)
					a *= 0.5
					currScale *= 2
				}
			} else {
				noise = noise2d(float64(x)*scale, float64(y)*scale)
			}
			noiseMap[x][y] = noise
		}
	}
	// normalize result
	normalizeNoiseArray(noiseMap)
	return noiseMap
}
