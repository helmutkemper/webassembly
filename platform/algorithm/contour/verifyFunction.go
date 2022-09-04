package contour

func (e *Contour) VerifyFunction(f func(pMatrix *[][]any, x, y int) bool) {
	e.verifyFunction = f
}
