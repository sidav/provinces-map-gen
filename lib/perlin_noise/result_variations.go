package perlin_noise

import (
	"math"
)

func normalizeNoiseArray(nm [][]float64) [][]float64 {
	min := 1000000000000.0
	max := -1000000000000.0
	for x := 0; x < len(nm); x++ {
		for y := 0; y < len(nm[x]); y++ {
			if min > nm[x][y] {
				min = nm[x][y]
			}
			if max < nm[x][y] {
				max = nm[x][y]
			}
		}
	}
	for x := 0; x < len(nm); x++ {
		for y := 0; y < len(nm[x]); y++ {
			nm[x][y] = (nm[x][y] - min) / (max - min)
		}
	}
	return nm
}

func MakeNoiseCircular(nm [][]float64) [][]float64 {
	w, h := len(nm), len(nm[0])
	centerX, centerY := w/2, h/2
	for x := 0; x < len(nm); x++ {
		for y := 0; y < len(nm[x]); y++ {
			grad := circleGrad(x, y, centerX, centerY, 15)
			nm[x][y] = nm[x][y] * grad
		}
	}
	normalizeNoiseArray(nm)
	return nm
}

func MakeNoiseRectangular(nm [][]float64) [][]float64 {
	w, h := len(nm), len(nm[0])
	centerX, centerY := w/2, h/2
	for x := 0; x < len(nm); x++ {
		for y := 0; y < len(nm[x]); y++ {
			grad := rectGrad(x, y, centerX, centerY, w/3, h/8, 15)
			nm[x][y] = nm[x][y] * grad
		}
	}
	normalizeNoiseArray(nm)
	return nm
}

func circleGrad(x, y, mapCenterX, mapCenterY, centerWeight int) float64 {
	dist := math.Sqrt(float64((x-mapCenterX)*(x-mapCenterX) + (y-mapCenterY)*(y-mapCenterY)))
	maxDist := math.Sqrt(float64(mapCenterX*mapCenterX + mapCenterY*mapCenterY))
	maxDist = float64(mapCenterX) * 2.5
	grad := dist / maxDist
	if grad > 1 {
		grad = 1
	}
	// make it from -1 to 1
	grad = (grad - 0.5) * -2
	if grad > 0.5 {
		grad = 0.5
	}
	if grad > 0 {
		grad *= float64(centerWeight)
	}
	return grad
}

func rectGrad(x, y, mapCenterX, mapCenterY, rectW, rectH, centerWeight int) float64 {
	dist := 0.0
	if !(math.Abs(float64(x-mapCenterX)) < float64(rectW/2) && math.Abs(float64(y-mapCenterY)) < float64(rectH)) {
		dist = math.Sqrt(float64(getSqDistFromCoordsToRectangleBorder(x, y, mapCenterX-rectW/2, mapCenterY-rectH/2, rectW, rectH)))
	}
	maxDist := math.Sqrt(float64(mapCenterX*mapCenterX + mapCenterY*mapCenterY))
	maxDist = float64(mapCenterY) * 2.5
	grad := dist / maxDist
	if grad > 1 {
		grad = 1
	}
	// make it from -1 to 1
	grad = (grad - 0.5) * -2
	if grad > 0.5 {
		grad = 0.5
	}
	if grad > 0 {
		grad *= float64(centerWeight)
	}
	return grad
}
