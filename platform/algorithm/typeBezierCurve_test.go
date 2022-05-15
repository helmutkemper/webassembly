package algorithm

import (
	"testing"
)

func TestBezierCurve_Density(t *testing.T) {
	var bezier = &BezierCurve{}
	bezier.Init()

	border := 50.0
	wight := 300.0
	height := 300.0

	// Português:
	// Pontos de um retângulo
	// E.g.: P0 (1,0) = (1*wight+border,0*height+border)
	// E.g.: P2 (2,1) = (2*wight+border,1*height+border)
	//
	//     (0,0)            (1,0)            (2,0)
	//       +----------------+----------------+
	//       | P7            P0             P1 |
	//       |                                 |
	//       |                                 |
	//       |                                 |
	// (0,1) + P6                           P2 + (2,1)
	//       |                                 |
	//       |                                 |
	//       |                                 |
	//       | P5            P4             P3 |
	//       +----------------+----------------+
	//     (0,2)            (1,2)            (2,2)

	bezier.Add(Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Add(Point{X: 2*wight + border, Y: 0*height + border})
	bezier.Add(Point{X: 2*wight + border, Y: 1*height + border})
	bezier.Add(Point{X: 2*wight + border, Y: 2*height + border})
	bezier.Add(Point{X: 1*wight + border, Y: 2*height + border})
	bezier.Add(Point{X: 0*wight + border, Y: 2*height + border})
	bezier.Add(Point{X: 0*wight + border, Y: 1*height + border})
	bezier.Add(Point{X: 0*wight + border, Y: 0*height + border})
	bezier.Add(Point{X: 1*wight + border, Y: 0*height + border})

	// Português:
	// Gera a curva de Bezié
	bezier.Process(0.001)

	// Reduz a curva a 5 segmentos de reta
	var decimatesCurve = &BezierCurve{}
	decimatesCurve.Copy(bezier)
	decimatesCurve.SetNumberOfSegments(3)

	var density = &BezierCurve{}
	density.Copy(decimatesCurve)
	density.IncreaseDensityBetweenPoints(10)

	okList := []Point{
		{X: 350, Y: 50},
		{X: 641, Y: 253},
		{X: 537, Y: 605},
		{X: 163, Y: 605},
		{X: 59, Y: 254},
		{X: 349, Y: 50},
	}
	for _, p := range decimatesCurve.processed {
		pass := false
		for _, pOk := range okList {

			// Português:
			// Arredonda os números decimais
			A := decimatesCurve.Round(p.X, 0) == pOk.X
			B := decimatesCurve.Round(p.Y, 0) == pOk.Y

			if A && B {
				pass = true
				break
			}
		}

		if pass == false {
			t.FailNow()
		}
	}

}

func TestBezierCurve(t *testing.T) {
	var bezier = &BezierCurve{}
	bezier.Init()

	border := 50.0
	wight := 350.0
	height := 350.0
	adjust := 0.0 //15.0

	// E.g.: P0 (1,0) = (1*wight,0*height)
	// E.g.: P2 (2,1) = (2*wight,1*height)
	//
	//     (0,0)            (1,0)            (2,0)
	//       +----------------+----------------+
	//       | P7            P0             P1 |
	//       |                                 |
	//       |                                 |
	//       |                                 |
	// (0,1) + P6                           P2 + (2,1)
	//       |                                 |
	//       |                                 |
	//       |                                 |
	//       | P5            P4             P3 |
	//       +----------------+----------------+
	//     (0,2)            (1,2)            (2,2)

	bezier.Add(Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Add(Point{X: 2*wight + border - adjust, Y: 0*height + border + adjust})
	bezier.Add(Point{X: 2*wight + border, Y: 1*height + border})
	bezier.Add(Point{X: 2*wight + border - adjust, Y: 2*height + border - adjust})
	bezier.Add(Point{X: 1*wight + border, Y: 2*height + border})
	bezier.Add(Point{X: 0*wight + border + adjust, Y: 2*height + border - adjust})
	bezier.Add(Point{X: 0*wight + border, Y: 1*height + border})
	bezier.Add(Point{X: 0*wight + border + adjust, Y: 0*height + border + adjust})
	bezier.Add(Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Process(0.001)

	var decimatesCurve = &BezierCurve{}
	decimatesCurve.Init()
	decimatesCurve.Copy(bezier)
	decimatesCurve.SetNumberOfSegments(5)

	var density = &BezierCurve{}
	density.Init()
	density.Copy(decimatesCurve)
	density.IncreaseDensityBetweenPoints(100)

	var ripple = &BezierCurve{}
	ripple.Init()
	ripple.Copy(bezier)
	ripple.GenerateRipple(30.0, 30)

}
