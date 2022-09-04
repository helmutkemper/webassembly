package contour

func (e *Contour) PopulateFunction(f func(pData *[][]any, x, y int)) {
	e.populateFunction = f
}
