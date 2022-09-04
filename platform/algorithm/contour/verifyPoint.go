package contour

func (e *Contour) verifyPoint(x, y int) {
	e.xStart = x
	e.yStart = y

	for {
		pass := false
		for i := -2; i <= pMax; i += 1 {

			p := e.pAdjust(e.pSpin + i)
			dx, dy := e.spin[p](x, y)
			if dx == 0 && dy == 0 {
				continue
			}
			e.pSpin = p

			e.populateFunction(&e.data, x, y)
			pass = true
			x += dx
			y += dy

			if x == e.xStart && y == e.yStart {
				return
			}

			break
		}

		if pass == false {
			break
		}
	}
}
