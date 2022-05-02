package algorithm

import (
	"testing"
)

func TestBezierCurve_Density(t *testing.T) {
	var curve BezierCurve
	curve.Init()

	b := 50.0
	dX := 500.0
	dY := 300.0

	//    0,0    1,0    2,0
	//     0------0------0
	//     |             |
	// 0,1 0             0 2,1
	//     |             |
	//     0------0------0
	//    0,2    1,2    2,2
	curve.Add(Point{X: 1*dX + b, Y: 0*dY + b})
	curve.Add(Point{X: 2*dX + b, Y: 0*dY + b})
	curve.Add(Point{X: 2*dX + b, Y: 1*dY + b})
	curve.Add(Point{X: 2*dX + b, Y: 2*dY + b})
	curve.Add(Point{X: 1*dX + b, Y: 2*dY + b})
	curve.Add(Point{X: 0*dX + b, Y: 2*dY + b})
	curve.Add(Point{X: 0*dX + b, Y: 1*dY + b})
	curve.Add(Point{X: 0*dX + b, Y: 0*dY + b})
	curve.Add(Point{X: 1*dX + b, Y: 0*dY + b})

	//curve.Process(0.1)

	curve.processed = append(curve.processed, Point{X: 0, Y: 0})   //0
	curve.processed = append(curve.processed, Point{X: 2, Y: 2})   //1
	curve.processed = append(curve.processed, Point{X: 4, Y: 4})   //2
	curve.processed = append(curve.processed, Point{X: 6, Y: 6})   //3
	curve.processed = append(curve.processed, Point{X: 7, Y: 7})   //4*
	curve.processed = append(curve.processed, Point{X: 8, Y: 8})   //5
	curve.processed = append(curve.processed, Point{X: 9, Y: 9})   //6*
	curve.processed = append(curve.processed, Point{X: 10, Y: 10}) //7
	curve.processed = append(curve.processed, Point{X: 11, Y: 11}) //8*
	curve.processed = append(curve.processed, Point{X: 11, Y: 11}) //9*
	curve.processed = append(curve.processed, Point{X: 11, Y: 11}) //10*
	curve.processed = append(curve.processed, Point{X: 11, Y: 11}) //11*
	curve.processed = append(curve.processed, Point{X: 11, Y: 11}) //12*
	curve.processed = append(curve.processed, Point{X: 12, Y: 12}) //13
	curve.processed = append(curve.processed, Point{X: 12, Y: 12}) //14*
	curve.processed = append(curve.processed, Point{X: 13, Y: 13}) //15*
	curve.processed = append(curve.processed, Point{X: 14, Y: 14}) //16
	curve.processed = append(curve.processed, Point{X: 15, Y: 15}) //17*
	curve.processed = append(curve.processed, Point{X: 15, Y: 15}) //18*
	curve.processed = append(curve.processed, Point{X: 16, Y: 16}) //19
	curve.processed = append(curve.processed, Point{X: 17, Y: 17}) //19

	curve.AdjustDensity()

	okList := []Point{{0, 0}, {2, 2}, {4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}, {15, 15}, {16, 16}, {17, 17}}
	for _, p := range curve.processed {
		pass := false
		for _, pOk := range okList {
			if p.X == pOk.X && p.Y == pOk.Y {
				pass = true
				break
			}
		}

		if pass == false {
			t.FailNow()
		}
	}

}
