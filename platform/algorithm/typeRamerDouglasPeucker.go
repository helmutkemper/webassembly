package algorithm

import (
	"math"
)

// Rdp
//
// English:
//
//  The Ramer–Douglas–Peucker algorithm, also known as the Douglas–Peucker algorithm and iterative
//  end-point fit algorithm, is an algorithm that decimates a curve composed of line segments to a
//  similar curve with fewer points.
//
// Português:
//
//  O algoritmo Ramer-Douglas-Peucker, também conhecido como algoritmo de Douglas-Peucker e algoritmo
//  iterativo de ajuste de ponto final, é um algoritmo que dizima uma curva composta de segmentos de
//  linha para uma curva semelhante com menos pontos.
type Rdp struct {
	original  []Point
	processed []Point
}

func (e *Rdp) Init() {
	e.original = make([]Point, 0)
	e.processed = make([]Point, 0)
}

func (e *Rdp) Add(p Point) {
	e.original = append(e.original, p)
}

func (e *Rdp) GetOriginal() (points []Point) {
	return e.original
}

func (e *Rdp) GetProcessed() (points []Point) {
	return e.processed
}

func (e *Rdp) Clear() {
	e.Init()
}

func (e *Rdp) ClearOriginal() {
	e.original = make([]Point, 0)
}

func (e *Rdp) ClearProcessed() {
	e.processed = make([]Point, 0)
}

func (e Rdp) findPerpendicularDistance(p, p1, p2 Point) (result float64) {
	if p1.X == p2.X {
		result = math.Abs(p.X - p1.X)
	} else {
		slope := (p2.Y - p1.Y) / (p2.X - p1.X)
		intercept := p1.Y - (slope * p1.X)
		result = math.Abs(slope*p.X-p.Y+intercept) / math.Sqrt(math.Pow(slope, 2)+1)
	}
	return
}

func (e *Rdp) Process(epsilon float64) {
	e.processed = e.subProcess(e.original, epsilon)
}

func (e *Rdp) subProcess(points []Point, epsilon float64) []Point {
	if len(points) < 3 {
		return points
	}
	firstPoint := points[0]
	lastPoint := points[len(points)-1]
	index := -1
	dist := float64(0)
	for i := 1; i < len(points)-1; i++ {
		cDist := e.findPerpendicularDistance(points[i], firstPoint, lastPoint)
		if cDist > dist {
			dist = cDist
			index = i
		}
	}
	if dist > epsilon {
		l1 := points[0 : index+1]
		l2 := points[index:]
		r1 := e.subProcess(l1, epsilon)
		r2 := e.subProcess(l2, epsilon)
		rs := append(r1[0:len(r1)-1], r2...)
		return rs
	} else {
		ret := make([]Point, 0)
		ret = append(ret, firstPoint, lastPoint)
		return ret
	}
}
