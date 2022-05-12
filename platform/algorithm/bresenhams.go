package algorithm

// Bresenham
//
// English:
//
//  Bresenham's line algorithm is a line drawing algorithm that determines the points of an
//  n-dimensional raster that should be selected in order to form a close approximation to a straight
//  line between two points.
//
// Português:
//
//  O algoritmo de linha de Bresenham é um algoritmo de desenho de linha que determina os pontos de um
//  raster n-dimensional que deve ser selecionado para formar uma aproximação próxima a uma linha reta
//  entre dois pontos.
type Bresenham struct {
	points []Point
}

func (e *Bresenham) Add(pos1, pos2 Point) {
	var points []Point

	x1, y1 := pos1.X, pos1.Y
	x2, y2 := pos2.X, pos2.Y

	isSteep := e.abs(y2-y1) > e.abs(x2-x1)
	if isSteep {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}

	reversed := false
	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
		reversed = true
	}

	deltaX := x2 - x1
	deltaY := e.abs(y2 - y1)
	err := deltaX / 2
	y := y1
	var yStep float64

	if y1 < y2 {
		yStep = 1
	} else {
		yStep = -1
	}

	for x := x1; x < x2+1; x++ {
		if isSteep {
			points = append(points, Point{X: y, Y: x})
		} else {
			points = append(points, Point{X: x, Y: y})
		}
		err -= deltaY
		if err < 0 {
			y += yStep
			err += deltaX
		}
	}

	if reversed {
		//Reverse the slice
		for i, j := 0, len(points)-1; i < j; i, j = i+1, j-1 {
			points[i], points[j] = points[j], points[i]
		}
	}

	e.points = append(e.points, points...)
	return
}

func (e Bresenham) Get() (list []Point) {
	return e.points
}

func (e Bresenham) abs(x float64) float64 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}
