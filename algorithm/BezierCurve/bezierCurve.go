package BezierCurve

type Point struct {
	X float64
	Y float64
}

type Curve struct {
	step  float64
	curve []Point
}

func (e Curve) getPercent(n1, n2 float64, percent float64) float64 {
	var diff = n2 - n1
	return n1 + (diff * percent)
}

func (e Curve) Get() (list []Point) {
	return e.curve
}

func (e Curve) Clear() {
	e.curve = make([]Point, 0)
}

func (e Curve) Add(x1, y1, x2, y2, x3, y3 float64) {
	var pX, pY, xA, xB, yA, yB float64

	if e.curve == nil {
		e.curve = make([]Point, 0)
	}

	if e.step == 0 {
		e.step = 0.01
	}

	for percent := 0.0; percent < 1.0; percent += e.step {
		xA = e.getPercent(x1, x2, percent)
		yA = e.getPercent(y1, y2, percent)
		xB = e.getPercent(x2, x3, percent)
		yB = e.getPercent(y2, y3, percent)

		pX = e.getPercent(xA, xB, percent)
		pY = e.getPercent(yA, yB, percent)

		e.curve = append(e.curve, Point{X: pX, Y: pY})
	}
}
