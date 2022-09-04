package contour

import "errors"

func (e *Contour) Init(pMatrix *[][]any) (err error) {

	if e.verifyFunction == nil {
		return errors.New("please, define the verify function before init")
	}

	if e.populateFunction == nil {
		return errors.New("please, define the populate function before init")
	}

	e.matrix = pMatrix

	e.yMax = len(*pMatrix)
	e.xMax = len((*pMatrix)[0])

	e.verified = make([][]bool, e.yMax)
	e.data = make([][]any, e.yMax)

	for y := 0; y != e.yMax; y += 1 {
		e.verified[y] = make([]bool, e.xMax)
		e.data[y] = make([]any, e.xMax)
	}

	e.yMax -= 1
	e.xMax -= 1

	e.spin = []walkingFunction{
		e.f0,   // 0
		e.f45,  // 1
		e.f90,  // 2
		e.f135, // 3
		e.f180, // 4
		e.f225, // 5
		e.f270, // 6
		e.f315, // 7
	}

	return
}
