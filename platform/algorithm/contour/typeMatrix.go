package contour

type Contour struct {
	verified [][]bool
	matrix   *[][]any
	data     [][]any
	x        int
	xMin     int
	xMax     int
	y        int
	yMin     int
	yMax     int

	xStart int
	yStart int

	verifyFunction   func(pMatrix *[][]any, x, y int) bool
	populateFunction func(pData *[][]any, x, y int)

	pSpin int
	spin  []walkingFunction
}
