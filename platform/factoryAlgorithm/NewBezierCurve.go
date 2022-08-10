package factoryAlgorithm

import "github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"

// NewBezierCurve
//
// The Bézier curve is a polynomial curve expressed as the linear interpolation between some representative points,
// called control points.
//
// Português:
//
// A curva de Bézier é uma curva polinomial expressa como a interpolação linear entre alguns pontos representativos,
// chamados pontos de controle.
func NewBezierCurve() (ref *algorithm.BezierCurve) {
	ref = new(algorithm.BezierCurve)
	ref.Init()

	return ref
}
