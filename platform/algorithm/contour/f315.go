package contour

func (e *Contour) f315(x, y int) (dx, dy int) {
	// x+1, y-1
	if x+1 <= e.xMax && y-1 >= e.yMin && e.verified[y-1][x+1] == false && e.verifyFunction(e.matrix, x+1, y-1) == true {
		return 1, -1
	}

	return 0, 0
}
