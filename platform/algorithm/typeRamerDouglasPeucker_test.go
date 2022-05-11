package algorithm

import (
	"testing"
)

func TestRamerDouglasPeucker(t *testing.T) {

	var rdp = &Rdp{}
	rdp.Init()

	rdp.Add(Point{0, 0})
	rdp.Add(Point{1, 1})
	rdp.Add(Point{2, 0})
	rdp.Add(Point{3, 2})
	rdp.Add(Point{4, 0})
	rdp.Add(Point{5, 3})
	rdp.Add(Point{6, 0})

	rdp.Process(2.0)
	for _, p := range *rdp.GetProcessed() {
		t.Logf("(%v,%v)", p.X, p.Y)
	}

}
