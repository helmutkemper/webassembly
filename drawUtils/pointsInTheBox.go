package drawUtils

import "math"

// PointsInTheBox receives a list of points drawn in a square box and calculates the new coordinates for when the box is resized
//
// Example:
//
//	Imagine de exclamation mark draw by the path:
//
//	M 185 120 L 190 240 L 210 240 L 215 120 L 185 120 z
//
//	M 190 260 L 190 280 L 210 280 L 210 260 L 190 260 z
//
//	points: [[185, 120], [190, 240], [210, 240], [215, 120], [185, 120],
//	         [190, 260], [190, 280], [210, 280], [210, 260], [190, 260]]
//
//	size from center: Math.min(width-2*margin, height-2*margin) / 2;
//
//	width and height are the new width and height of the box
func PointsInTheBox(points [][]int, size, width, height int, rotation float64) [][]int {
	const infinity = math.MaxInt
	minX, maxX := infinity, -infinity
	minY, maxY := infinity, -infinity

	// Determinar os limites da forma
	for _, point := range points {
		x, y := point[0], point[1]
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	shapeWidth := maxX - minX
	shapeHeight := maxY - minY

	// Escala baseada no menor lado da forma ajustada ao tamanho
	scale := float64(size) / float64(max(shapeWidth, shapeHeight))

	// Centro do canvas
	centerX := float64(width) / 2.0
	centerY := float64(height) / 2.0

	cosA := math.Cos(rotation)
	sinA := math.Sin(rotation)

	pointsCalculated := make([][]int, 0, len(points))

	// Recalcular os pontos
	for _, point := range points {
		x, y := point[0], point[1]

		// Normalizar em relação ao centro da forma
		normX := (float64(x) - (float64(minX) + float64(shapeWidth)/2.0)) * scale
		normY := (float64(y) - (float64(minY) + float64(shapeHeight)/2.0)) * scale

		// Aplicar rotação
		rotX := normX*cosA - normY*sinA
		rotY := normX*sinA + normY*cosA

		// Deslocar para o centro do canvas
		px := centerX + rotX
		py := centerY + rotY

		pointsCalculated = append(pointsCalculated, []int{int(px), int(py)})
	}

	return pointsCalculated
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
