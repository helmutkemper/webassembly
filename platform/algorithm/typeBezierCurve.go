package algorithm

// BezierCurve
//
//  The Bézier curve is a polynomial curve expressed as the linear interpolation between some
//  representative points, called control points.
//
// Português:
//
//  A curva de Bézier é uma curva polinomial expressa como a interpolação linear entre alguns pontos
//  representativos, chamados pontos de controle.
type BezierCurve struct {
	step      float64
	original  []Point
	processed []Point
}

func (e *BezierCurve) getPercent(n1, n2 float64, percent float64) float64 {
	var diff = n2 - n1
	return n1 + (diff * percent)
}

func (e *BezierCurve) GetProcessed() (list *[]Point) {
	return &e.processed
}

func (e *BezierCurve) GetOriginal() (list *[]Point) {
	return &e.original
}

func (e *BezierCurve) ClearProcessed() {
	e.processed = make([]Point, 0)
}

func (e *BezierCurve) ClearOriginal() {
	e.original = make([]Point, 0)
}

func (e *BezierCurve) Clear() {
	e.Init()
}

func (e *BezierCurve) Init() {
	e.processed = make([]Point, 0)
	e.original = make([]Point, 0)
}

func (e *BezierCurve) Add(p Point) {
	e.original = append(e.original, p)
}

func (e *BezierCurve) Process(step float64) {
	e.step = step
	if len(e.original) < 3 {
		e.processed = make([]Point, len(e.original))
		copy(e.processed, e.original)
		return
	}

	var p1, p2, p3 Point
	var l = len(e.original)
	for i := 0; i < l; i += 2 {

		if i+2 >= l {
			break
		}

		p1 = e.original[i+0]
		p2 = e.original[i+1]
		p3 = e.original[i+2]
		e.subProcess(p1, p2, p3, step)
	}
}

func (e *BezierCurve) subProcess(p1 Point, p2 Point, p3 Point, step float64) {
	var pX, pY, xA, xB, yA, yB float64

	for percent := 0.0; percent < 1.0; percent += e.step {
		xA = e.getPercent(p1.X, p2.X, percent)
		yA = e.getPercent(p1.Y, p2.Y, percent)
		xB = e.getPercent(p2.X, p3.X, percent)
		yB = e.getPercent(p2.Y, p3.Y, percent)

		pX = e.getPercent(xA, xB, percent)
		pY = e.getPercent(yA, yB, percent)

		e.processed = append(e.processed, Point{X: pX, Y: pY})
	}
}
