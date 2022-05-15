package algorithm

import (
	"math"
)

type density struct {
	Geometry
}

// increaseDensityBetweenPoints
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
func (e *density) increaseDensityBetweenPoints(lineSegments int, processed *[]Point) {

	if len(*processed) == 0 {
		return
	}

	tmp := make([]Point, len(*processed))
	copy(tmp, *processed)

	*processed = make([]Point, 0)
	p1 := Point{}
	p2 := Point{}

	length := len(tmp) - 1
	for k := range tmp {
		if k != length {
			p1 = (tmp)[k]
			p2 = (tmp)[k+1]

			distance := math.Sqrt(math.Pow(p1.X-p2.X, 2.0) + math.Pow(p1.Y-p2.Y, 2.0))
			angle := math.Atan2(p1.Y-p2.Y, p1.X-p2.X)
			cos := math.Cos(angle)
			sin := math.Sin(angle)

			*processed = append(
				*processed,
				p1,
			)

			for i := 0; i != lineSegments-1; i += 1 {

				percent := 1.0 / float64(lineSegments) * (1.0 + float64(i))
				x := p1.X - cos*distance*percent
				y := p1.Y - sin*distance*percent
				pNew := Point{X: x, Y: y}

				*processed = append(
					*processed,
					pNew,
				)
			}
		}
	}

	*processed = append(
		*processed,
		p2,
	)
}

// setNumberOfSegments
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
func (e *density) setNumberOfSegments(n int, original, processed *[]Point) {
	//*processed = make([]Point, len(*original))
	//copy(*processed, *original)

	tmp := make([]Point, 0)

	d := 0.0
	l := len(*processed) - 1
	for i := 0; i != l; i += 1 {
		dCalc := e.DistanceBetweenTwoPoints((*processed)[i], (*processed)[i+1])
		d += dCalc
		(*processed)[i+1].d = d
	}

	dTarget := d / float64(n)
	p3 := Point{}

	var dList = make([]float64, 0)

	for i := 0; i != n; i += 1 {
		dList = append(dList, dTarget*float64(i))
	}

	dActual := 0.0
	dTargetTmp := dTarget
	for i := 0; i != l; i += 1 {
		dActual = (*processed)[i].d
		dTargetTmp = dList[0] - dActual
		if dActual > dList[0] {
			dCalc := e.DistanceBetweenTwoPoints((*processed)[i-1], (*processed)[i]) - (-1 * dTargetTmp)
			p3 = e.PointBetweenTwoPoints((*processed)[i-1], (*processed)[i], dCalc)
			tmp = append(tmp, p3)

			if len(dList) == 1 {
				break
			}

			dList = dList[1:]
		} else if dActual == dList[0] {
			p3 = (*processed)[i]
			tmp = append(tmp, p3)

			if len(dList) == 1 {
				break
			}

			dList = dList[1:]
		}
	}

	tmp = append(tmp, (*processed)[len(*processed)-1])

	*processed = make([]Point, len(tmp))
	copy(*processed, tmp)
}
