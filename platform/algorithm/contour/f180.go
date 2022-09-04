package contour

func (e *Contour) f180(x, y int) (dx, dy int) {
	// x-1
	if x-1 >= e.xMin && e.verified[y][x-1] == false && e.verifyFunction(e.matrix, x-1, y) == true {
		return -1, 0
	}

	return 0, 0
}
