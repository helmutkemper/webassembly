package algorithm

import (
	"log"
	"testing"
)

func TestBezierCurve_Density(t *testing.T) {
	var curve = &Density{}
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
	curve.processed = append(curve.processed, Point{X: 10, Y: 10}) //7
	curve.processed = append(curve.processed, Point{X: 11, Y: 11}) //8*
	curve.processed = append(curve.processed, Point{X: 11, Y: 11}) //9*
	curve.processed = append(curve.processed, Point{X: 11, Y: 11}) //10*
	curve.processed = append(curve.processed, Point{X: 13, Y: 13}) //15*
	curve.processed = append(curve.processed, Point{X: 14, Y: 14}) //16
	curve.processed = append(curve.processed, Point{X: 15, Y: 15}) //17*
	curve.processed = append(curve.processed, Point{X: 15, Y: 15}) //18*
	curve.processed = append(curve.processed, Point{X: 16, Y: 16}) //19
	curve.processed = append(curve.processed, Point{X: 17, Y: 17}) //19

	curve.AdjustDensityByNSegments(3)

	for k, p := range curve.processed {
		if k == 0 {
			log.Printf("(%v,%v)", p.X, p.Y)
			continue
		}
		p1 := curve.processed[k-1]
		p2 := curve.processed[k]
		log.Printf("(%v,%v) :%v", p.X, p.Y, curve.Round(curve.Distance(p1, p2), 2))
	}

	okList := []Point{{0, 0}, {1.5, 1.5}, {3.0, 3.0}, {4.5, 4.5}, {6.0, 6.0},
		{7.5, 7.5}, {9.0, 9.0}, {10.5, 10.5}, {12.0, 12.0}, {13.5, 13.5},
		{15.0, 15.0}, {16.5, 16.5}, {17.0, 17.0}}
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
