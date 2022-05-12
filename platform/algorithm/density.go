package algorithm

import (
	"math"
)

type Density struct {
	original  []Point
	processed []Point
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
func (e *Density) IncreaseDensityBetweenPoints(lineSegments int) {

	e.ClearProcessed()

	if len(e.original) == 0 {
		return
	}

	length := len(e.original) - 1
	for k := range e.original {
		if k != length {
			p1 := e.original[k]
			p2 := e.original[k+1]

			e.processed = append(
				e.processed,
				p1,
			)

			for i := 0; i != lineSegments-1; i += 1 {
				div := 1.0 / float64(lineSegments) * (1.0 + float64(i))
				d := math.Sqrt(math.Pow(p1.X-p2.X, 2.0)+math.Pow(p1.Y-p2.Y, 2.0)) * div
				a := math.Atan2(p1.Y-p2.Y, p1.X-p2.X)
				x := p1.X - math.Cos(a)*d
				y := p1.Y - math.Sin(a)*d
				pNew := Point{X: x, Y: y}

				e.processed = append(
					e.processed,
					pNew,
				)

			}

			e.processed = append(
				e.processed,
				p2,
			)
		}
	}
}

// PointBetweenTwoPoints
//
// English:
//
//  Returns a point between two points from the distance of P1.
//
//   Input:
//     p1, p2: start and end point of the line;
//     distance: distance of the new point from p1.
//
//   Output:
//     p: midpoint between p1 and p2.
//
// Português:
//
//  Devolve um ponto entre dois pontos a partir da distância de P1.
//
//   Entrada:
//     p1, p2: ponto inicial e final da reta;
//     distance: distância do novo ponto em relação a p1.
//
//   Saída:
//     p: ponto intermediário entre p1 e p2.
func (e *Density) PointBetweenTwoPoints(p1, p2 Point, distance float64) (p Point) {
	a := math.Atan2(p1.Y-p2.Y, p1.X-p2.X)
	x := e.Round(p1.X-math.Cos(a)*distance, 2)
	y := e.Round(p1.Y-math.Sin(a)*distance, 2)
	p = Point{X: x, Y: y}
	return
}

// Round
//
// English:
//
//
//
//
//
// Português:
//
//  Arredonda a quantidade máxima de casas decimais de um número de ponto flutuante.
//
//   Entrada:
//     value: número a ser arredondado;
//     places: quantidade de dígitos decimais.
//
//   Notes:
//     * Se o dígito seguinte a places for menor que 5, o dígito place será arredondado para baixo.
//       Exemplo: Round(3.1415, 2) = 3.14
//                Round(3.1415, 3) = 3.142
func (e *Density) Round(value float64, places int) float64 {
	var roundOn = 0.5

	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * value
	_, div := math.Modf(digit)

	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	return round / pow
}

// Distance
//
// English:
//
//  Calculates the distance between two points.
//
//   Input:
//     p1, p2: Start and end points of the line.
//
//   Output:
//     d: distance between p1 and p2 in pixels.
//
// Português:
//
//  Calcula a distância entre dois pontos.
//
//   Entrada:
//     p1, p2: Pontos inicial e final da reta.
//
//   Saída:
//     d: distância entre p1 e p2 em pixels.
func (e *Density) Distance(p1, p2 Point) (d float64) {
	return math.Sqrt(math.Abs(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2)))
}

type Mode int

const (
	KModeMin = iota
	KModeMinR
	KModeMed
	KModeMax
)

func (e *Density) AdjustDensityByDistance(distance float64) (ref *Density) {

	newList := make([]Point, 0)
	newList = append(newList, e.processed[0])
	l := len(e.processed) - 1

	var p3 Point
	var pTest int
	var lastPoint = e.processed[0]
	for pInitial := 0; pInitial < l; {
		for pTest = pInitial + 1; pTest < l+1; pTest += 1 {
			d := e.Distance(lastPoint, e.processed[pTest])
			if d >= distance {
				if d > distance {
					p3 = e.PointBetweenTwoPoints(lastPoint, e.processed[pTest], distance)
				} else {
					p3 = e.processed[pTest]
				}
				newList = append(newList, p3)
				lastPoint = p3
				break
			}
		}

		pInitial = pTest - 1
	}

	p1 := e.processed[len(e.processed)-1]
	p2 := newList[len(newList)-1]

	if p1.X != p2.X || p1.Y != p2.Y {
		newList = append(newList, p1)
	}

	e.processed = make([]Point, len(newList))
	copy(e.processed, newList)

	return e
}

func (e *Density) AdjustDensity(mode Mode) (ref *Density) {
	dSelected := 0.0
	dMax := (math.MaxFloat64 - 1) * -1
	dMin := math.MaxFloat64

	l := len(e.processed) - 1
	for i := 0; i != l; i += 1 {
		d := e.Distance(e.processed[i], e.processed[i+1])
		dMax = math.Max(d, dMax)
		dTmp := math.Min(d, dMin)
		if mode != KModeMinR && dTmp != 0 {
			dMin = dTmp
		}
	}

	switch mode {
	case KModeMin:
		dSelected = dMin
	case KModeMinR:
		dSelected = dMin
	case KModeMed:
		dSelected = (dMax + dMin) / 2
	case KModeMax:
		dSelected = dMax
	}

	e.AdjustDensityByDistance(dSelected)

	return e
}

// 2,5
//
// 0,0                         3,4                         6,8
//  +-------------+-------------+-------------+-------------+
//
// 0,0         1.5,2.0       3.0,4.0       4.5,6.0       6.0,8.0
//  +-------------+-------------+-------------+-------------+
//  |          dPoint           |
//                |    dt-dp    |
//
//
// 0,0         1.5,2.0
//  +-------------*-----------------------------------------+
//
// d:5.0
// t:2.5
//
//
//
//
//
//
//

func (e *Density) AdjustDensityByNSegments(n int) (ref *Density) {
	e.processed = make([]Point, len(e.original))
	copy(e.processed, e.original)

	tmp := make([]Point, 0)
	tmp = append(tmp, e.original[0])

	d := 0.0
	l := len(e.processed) - 1
	for i := 0; i != l; i += 1 {
		dCalc := e.Distance(e.processed[i], e.processed[i+1])
		e.processed[i+1].d = dCalc
		d += dCalc
	}

	dTarget := d / float64(n)
	//lastPoint := e.original[0]
	p3 := Point{}

	dActual := 0.0
	dTargetTmp := dTarget
	for i := 0; i != l; i += 1 {
		dActual += e.processed[i+1].d
		if dActual > dTargetTmp {
			p3 = e.PointBetweenTwoPoints(e.processed[i], e.processed[i+1], dTarget)
			dTargetTmp = dActual - dTargetTmp
			tmp = append(tmp, p3)
		} else if dActual == dTargetTmp {
			p3 = e.PointBetweenTwoPoints(e.processed[i], e.processed[i+1], dTarget)
			dTargetTmp = dTarget
			tmp = append(tmp, p3)
		} else {
			dTargetTmp -= dActual
		}
	}

	//for i := 0; i != l; i += 1 {
	//	d = e.Distance(lastPoint, e.processed[i+1])
	//	if dActual+d > dTargetTmp {
	//		dCalc := dActual + d - dTargetTmp
	//		p3 = e.PointBetweenTwoPoints(e.processed[i], e.processed[i+1], dCalc)
	//		lastPoint = p3
	//		dActual = d - dTargetTmp + dActual
	//		tmp = append(tmp, p3)
	//	} else if dActual+d == dTarget {
	//		p3 = e.processed[i+1]
	//		lastPoint = p3
	//		dActual = 0.0
	//		tmp = append(tmp, p3)
	//	} else {
	//		lastPoint = e.processed[i]
	//		dActual += d
	//	}
	//
	//}
	tmp = append(tmp, e.processed[len(e.processed)-1])

	e.processed = make([]Point, len(tmp))
	copy(e.processed, tmp)

	return e
}

func (e *Density) Add(p Point) (ref *Density) {
	e.original = append(e.original, p)

	return e
}

func (e *Density) GetProcessed() (list *[]Point) {
	return &e.processed
}

func (e *Density) GetOriginal() (list *[]Point) {
	return &e.original
}

func (e *Density) ClearProcessed() (ref *Density) {
	e.processed = make([]Point, 0)

	return e
}

func (e *Density) ClearOriginal() (ref *Density) {
	e.original = make([]Point, 0)

	return e
}

func (e *Density) Clear() (ref *Density) {
	e.Init()

	return e
}

func (e *Density) Init() (ref *Density) {
	e.processed = make([]Point, 0)
	e.original = make([]Point, 0)

	return e
}
