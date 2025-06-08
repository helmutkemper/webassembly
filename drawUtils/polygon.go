package drawUtils

import "math"

// Polygon generates the points of a regular polygon given the number of sides, radius, width, height and rotation
func Polygon(sides int, radius, width, height int, rotation float64) [][]int {
	// Centro do canvas
	cx := float64(width) / 2.0
	cy := float64(height) / 2.0

	// Gerar os pontos do polígono
	points := make([][]int, 0, sides+1)
	for i := 0; i <= sides; i++ {
		angle := (2*math.Pi*float64(i))/float64(sides) + rotation
		x := cx + float64(radius)*math.Cos(angle)
		y := cy + float64(radius)*math.Sin(angle)
		points = append(points, []int{int(x), int(y)})
	}

	return points
}
