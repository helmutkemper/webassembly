package contour

func (e *Contour) Vector(x, y int, degrees Degrees) {
	e.pSpin = int(degrees)
	e.verifyPoint(x, y)
}
