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
	Geometry
	ripple  *ripple
	density density
	step    float64
}

func (e *BezierCurve) GetProcessed() (list *[]Point) {
	return &e.processed
}

func (e *BezierCurve) GetOriginal() (list *[]Point) {
	return &e.original
}

func (e *BezierCurve) Copy(ref CopyInterface) {
	e.original = make([]Point, len(*ref.GetOriginal()))
	e.processed = make([]Point, len(*ref.GetProcessed()))

	copy(e.original, *ref.GetOriginal())
	copy(e.processed, *ref.GetProcessed())
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

	//todo: make fabric
	e.ripple = &ripple{}
}

func (e *BezierCurve) Add(point Point) {
	e.original = append(e.original, point)
}

func (e *BezierCurve) getPercent(n1, n2 float64, percent float64) float64 {
	var diff = n2 - n1
	return n1 + (diff * percent)
}

func (e *BezierCurve) Process(step float64) (ref *BezierCurve) {
	e.step = step
	if len(e.original) < 3 {
		e.processed = make([]Point, len(e.original))
		copy(e.processed, e.original)

		return e
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
		e.subProcess(p1, p2, p3)
	}

	return e
}

func (e *BezierCurve) subProcess(p1 Point, p2 Point, p3 Point) {
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

// GenerateRipple
//
// English:
//
//  Generates a sine ripple in the list of processed points, losing the original points
//
//   Input:
//     distance: maximum ripple distance, wavelength;
//     ripples: amounts of ripples;
//
// Português:
//
//  Gera uma ondulação senoidal na lista de pontos processados, perdendo os pontos originais
//
//   Entrada:
//     distance: distância máxima da ondulação, ou comprimento de onda;
//     ripples: quantidades de ondulações;
//     processed: ponteiro da lista de pontos processados.
func (e *BezierCurve) GenerateRipple(distance float64, ripples int) {
	e.ripple.generateRipple(distance, ripples, &e.processed)
}

// IncreaseDensityBetweenPoints
//
// English:
//
//  Increases the number of segments between points, in a straight line.
//
//   Input:
//     lineSegments: number of line segments to be added between points.
//
//   Notes:
//     * Use this function when the stitch density is low when using easing tween functions, so the movement is more
//       fluid.
//
// Português:
//
//  Aumenta a quantidade de segmentos entre os pontos, em linha reta.
//
//   Entrada:
//     lineSegments: quantidade de seguimentos de reta a serem adicionados entre os pontos.
//
//   Nota:
//     * Use esta função quando a densidade de pontos for baixa durante o uso das funções easing tween, para que o
//       movimento fique mais fluido.
func (e *BezierCurve) IncreaseDensityBetweenPoints(lineSegments int) {
	e.density.increaseDensityBetweenPoints(lineSegments, &e.processed)
}

// SetNumberOfSegments
//
// English:
//
//  Adjusts the number of line segments that make up the curve.
//
// Every curve is formed by N straight segments, where a curve with many segments of irregular sizes makes the movement
// of objects unstable, so adjusting the number of segments when generating a curve makes the movement more uniform,
// in addition to saving memory.
//
// Português:
//
//  Ajusta a quantidade de segmentos de reta que formam a curva.
//
// Toda curva é formada N segmentos de reta, onde uma curva com muitos seguimentos de tamanhos irregulares deixam o
// movimento de objetos instável, por isto, ajustar o número de segmentos ao gerar uma curva, deixa o movimento mais
// uniforme, além de poupar memória.
func (e *BezierCurve) SetNumberOfSegments(n int) {
	e.density.setNumberOfSegments(n, &e.original, &e.processed)
}
