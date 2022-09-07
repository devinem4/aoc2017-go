package utils

import "math"

func CalcManhattanDist(x int, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}