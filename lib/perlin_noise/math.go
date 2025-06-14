package perlin_noise

type vector struct {
	x, y float64
}

func (v *vector) dotProduct(other *vector) float64 {
	return v.x*other.x + v.y*other.y
}

func maxInt(args ...int) int {
	currMax := 0
	for i, arg := range args {
		if i == 0 || arg > currMax {
			currMax = arg
		}
	}
	return currMax
}

func getSqDistFromCoordsToRectangleBorder(x, y, rx, ry, w, h int) int {
	//dx := maxInt(rx - x, 0, x - (rx+w-1))
	//dy := maxInt(ry - y, 0, y - (ry+h-1))
	dx := maxInt(rx-x, x-(rx+w-1))
	dy := maxInt(ry-y, y-(ry+h-1))
	return dx*dx + dy*dy
}
