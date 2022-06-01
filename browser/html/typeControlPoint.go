package html

import (
	"fmt"
)

// ControlPoint
//
// English:
//
// Each control point description is a set of four values: x1 y1 x2 y2, describing the Bézier control points for one
// time segment. The keyTimes values that define the associated segment are the Bézier "anchor points", and the
// keySplines values are the control points. Thus, there must be one fewer sets of control points than there are
// keyTimes.
//
// The values of x1 y1 x2 y2 must all be in the range 0 to 1.
//
// Português:
//
// Cada descrição de ponto de controle é um conjunto de quatro valores: x1 y1 x2 y2, descrevendo os pontos de controle
// Bézier para um segmento de tempo. Os valores de keyTimes que definem o segmento associado são os "pontos de
// ancoragem" de Bézier, e os valores de keySplines são os pontos de controle. Assim, deve haver um conjunto de pontos
// de controle a menos do que keyTimes.
//
// Os valores de x1 y1 x2 y2 devem estar todos no intervalo de 0 a 1.
type ControlPoint struct {
	l []cp
}

func (e *ControlPoint) Add(x1, y1, x2, y2 float64) {
	if e.l == nil {
		e.l = make([]cp, 0)
	}

	e.l = append(e.l, cp{X1: x1, Y1: y1, X2: x2, Y2: y2})
}

func (e ControlPoint) String() string {
	ret := ""
	for _, point := range e.l {
		ret += point.String()
	}

	return ret
}

type cp struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func (e cp) String() string {
	return fmt.Sprintf("%v %v %v %v;", e.X1, e.Y1, e.X2, e.Y2)
}
