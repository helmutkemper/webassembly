package main

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/factoryFontFamily"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
	"github.com/helmutkemper/webassembly/textUtil"
	"image/color"
	"math"
	"syscall/js"
)

type Point struct {
	X, Y float64
}

type Circle struct {
}

func (e Circle) PointInside(pt, c Point, r float64) bool {
	// d = √((Xp - Xc)² + (Yp - Yc)²)
	d := math.Sqrt(math.Pow(pt.X-c.X, 2) + math.Pow(pt.Y-c.Y, 2))
	return d < r
}

type Vertex struct {
	v1 Point
	v2 Point
	v3 Point
}

func (e *Vertex) Set(c, p1, p2 Point) {
	e.v1.X = c.X
	e.v1.Y = c.Y

	e.v2.X = p1.X
	e.v2.Y = p1.Y

	e.v3.X = p2.X
	e.v3.Y = p2.Y
}

func (e *Vertex) GetXY() (p1, p2 Point) {
	return e.v2, e.v3
}

type Triangle struct {
	v1, v2, v3 Point
}

func (e Triangle) sign(p1, p2, p3 Point) float64 {
	return (p1.X-p3.X)*(p2.Y-p3.Y) - (p2.X-p3.X)*(p1.Y-p3.Y)
}

func (e Triangle) PointInside(pt Point, v Vertex) bool {
	d1 := e.sign(pt, v.v1, v.v2)
	d2 := e.sign(pt, v.v2, v.v3)
	d3 := e.sign(pt, v.v3, v.v1)

	hasNeg := d1 < 0 || d2 < 0 || d3 < 0
	hasPos := d1 > 0 || d2 > 0 || d3 > 0

	return !(hasNeg && hasPos)
}

type Polygon struct {
	objTriangle Triangle
	box         Box
	triangles   []Vertex
	col         int
	row         int
	cx          int
	cy          int
}

func (e *Polygon) PointInside(pt Point) bool {
	if !e.box.PointInside(pt) {
		return false
	}

	for _, t := range e.triangles {
		if e.objTriangle.PointInside(pt, t) {
			return true
		}
	}

	return false
}

func (e *Polygon) GetVertex() (v []Vertex) {
	return e.triangles
}

func (e *Polygon) SetVertex(c, p1, p2 Point) {
	v := Vertex{}
	v.Set(c, p1, p2)

	if e.triangles == nil {
		e.triangles = make([]Vertex, 0)
	}

	e.triangles = append(e.triangles, v)
}

func (e *Polygon) SetAddr(col, row int) {
	e.col = col
	e.row = row
}

func (e *Polygon) GetAddr() (col, row int) {
	return e.col, e.row
}

func (e *Polygon) SetCenter(x, y int) {
	e.cx = x
	e.cy = y
}

func (e *Polygon) GetCenter() (x, y int) {
	return e.cx, e.cy
}

func (e *Polygon) SetBox(c Point, r float64) {
	e.box.Set(c, r)
}

type Box struct {
	topLeft     Point
	bottomRight Point
}

func (e *Box) Set(c Point, r float64) {
	e.topLeft.X = c.X - r
	e.topLeft.Y = c.Y - r

	e.bottomRight.X = c.X + r
	e.bottomRight.Y = c.Y + r
}

func (e *Box) PointInside(pt Point) bool {
	if !(e.topLeft.X <= pt.X && e.bottomRight.X >= pt.X) {
		return false
	}

	if !(e.topLeft.Y <= pt.Y && e.bottomRight.Y >= pt.Y) {
		return false
	}
	return true
}

type HexagonDraw struct {
	svg            *html.TagSvg
	polygon        []Polygon
	currentPolygon *Polygon
	rotation       float64
	sides          int
	space          int
	radius         int
}

func (e *HexagonDraw) Init() {
	e.rotation = 0.0
	e.sides = 6
	e.space = 10
	e.radius = 100
	e.polygon = make([]Polygon, 0)

	e.svg = factoryBrowser.NewTagSvg().
		Width("100vw").
		Height("100vh").
		Append(
			factoryBrowser.NewTagSvgFilter().Id("blur1").Append(
				factoryBrowser.NewTagSvgFeOffset().Dx(4).Dy(3),
				factoryBrowser.NewTagSvgFeBlend().In2("offOut"),
				factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(8),
			),
			factoryBrowser.NewTagSvgFilter().Id("blur2").Append(
				factoryBrowser.NewTagSvgFeOffset().Dx(2).Dy(2),
				factoryBrowser.NewTagSvgFeBlend().In2("offOut"),
				factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(2),
			),
			factoryBrowser.NewTagSvgFilter().Id("blur3").Append(
				factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(0.5),
			),
		)
}

// calculatePolygon generates the points of a regular polygon given the number of sides, radius, width, height and rotation
func (e *HexagonDraw) calculatePolygon(sides, radius, cx, cy int, rotation float64) [][]int {
	// Centro do canvas

	// Gerar os pontos do polígono
	points := make([][]int, 0, sides+1)
	for i := 0; i <= sides; i++ {
		angle := (2*math.Pi*float64(i))/float64(sides) + rotation
		x := float64(cx) + float64(radius)*math.Sin(angle)
		y := float64(cy) + float64(radius)*math.Cos(angle)
		points = append(points, []int{mathUtil.FloatToInt(x), mathUtil.FloatToInt(y)})

		if i != 0 {
			e.currentPolygon.SetVertex(
				Point{X: float64(cx), Y: float64(cy)},
				Point{X: x, Y: y},
				Point{X: float64(points[i-1][0]), Y: float64(points[i-1][1])},
			)
		}
	}

	e.polygon = append(e.polygon, *e.currentPolygon)
	return points
}

func (e *HexagonDraw) Point() {

}

func (e *HexagonDraw) CenterToXY(row, col, r int) (x int, y int) {

	//               . . . . . . .             --+--
	//             .       r        .            |
	//          .r                    r.
	//       .                            .
	//    .                                 .    h   1
	//    .                                 .    e   .
	//    .r               .               r.    i   5
	//    .               . .               .    g   *
	//    .              .   .              .    h   r
	//       .  r      r.     .r      r  .       t
	//          .      .       .      .
	//             .  .    r    .  .
	//               . . . . . . .               |
	//    |-          width = √3*r          -| --+--

	w := float64(r) * math.Sqrt(3)
	h := float64(r) * 1.5
	cx := float64(col) * w

	// approaches the left border
	cx -= w / 2.0

	if row%2 == 1 {
		cx += w / 2
	}

	cy := float64(row) * h

	// approaches the top
	cy -= 1.0 / 4.0 * h

	x = mathUtil.FloatToInt(cx)
	y = mathUtil.FloatToInt(cy)

	return
}

func (e *HexagonDraw) Make(col, row int) {
	x, y := e.CenterToXY(row, col, e.radius)

	e.currentPolygon = new(Polygon)
	e.currentPolygon.SetAddr(col, row)
	e.currentPolygon.SetBox(Point{X: float64(x), Y: float64(y)}, float64(e.radius))
	e.currentPolygon.SetCenter(x, y)

	e.calculatePolygon(e.sides, e.radius-e.space, x, y, e.rotation)
}

func (e *HexagonDraw) Draw() *HexagonDraw {
	for _, polygon := range e.polygon {
		hexagonExternalPath := make([]string, 0)

		for k, vertex := range polygon.triangles {
			p1, _ := vertex.GetXY()

			switch k {
			case 0:
				hexagonExternalPath = append(
					hexagonExternalPath,
					fmt.Sprintf("M %v %v", p1.X, p1.Y), // Move to the first point
				)
			default:
				hexagonExternalPath = append(hexagonExternalPath, fmt.Sprintf("L %v %v", p1.X, p1.Y)) // Draw lines to remaining points
			}
		}

		hexagonExternalPath = append(hexagonExternalPath, "z") // Close the path

		g := factoryBrowser.NewTagSvgG()

		path := factoryBrowser.NewTagSvgPath().
			Fill("none").
			Stroke("black").
			StrokeWidth(1).
			Filter("url(#blur2)")
		path.D(hexagonExternalPath)
		g.Append(path)

		path = factoryBrowser.NewTagSvgPath().
			Fill("none").
			Stroke("black").
			StrokeWidth(1).
			Filter("url(#blur1)")
		path.D(hexagonExternalPath)
		g.Append(path)

		path = factoryBrowser.NewTagSvgPath().
			Fill("white").
			Stroke(color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			}).
			StrokeWidth(2).
			Filter("url(#blur3)")
		path.D(hexagonExternalPath)
		g.Append(path)

		t := fmt.Sprintf("C: %v, L: %v", polygon.col, polygon.row)
		font := factoryFontFamily.NewArialBlack()
		fontSize := 30
		w, h := textUtil.GetTextSize(
			t,
			font,
			false,
			false,
			fontSize,
		)
		text := factoryBrowser.NewTagSvgText().
			X(polygon.cx - w/2).
			Y(polygon.cy + h/3).
			FontSize(fontSize).
			FontFamily(font).
			Text(t)
		g.Append(text)

		e.svg.Append(g)
	}

	return e
}

func (e *HexagonDraw) ReDraw(cx, cy, r int) {
	circle := factoryBrowser.NewTagSvgCircle().
		Cx(cx).
		Cy(cy).
		R(r).
		Fill("none").
		Stroke("red").
		StrokeWidth(1)
	e.svg.Append(circle)
}

func (e *HexagonDraw) Test(x, y int) {

	for _, polygon := range e.polygon {
		if polygon.PointInside(Point{float64(x), float64(y)}) {
			cx, cy := polygon.GetCenter()
			e.ReDraw(cx, cy, 20)
			vList := polygon.GetVertex()

			for _, v := range vList {
				p := []string{
					fmt.Sprintf("M %v %v", int(v.v1.X), int(v.v1.Y)),
					fmt.Sprintf("L %v %v", int(v.v2.X), int(v.v2.Y)),
					fmt.Sprintf("L %v %v", int(v.v3.X), int(v.v3.Y)),
				}

				path := factoryBrowser.NewTagSvgPath().
					Fill("none").
					Stroke("red").
					StrokeWidth(1).
					D(p)
				e.svg.Append(path)
			}
		}

	}
}

func main() {

	h := new(HexagonDraw)
	h.Init()

	table := [][]int{

		{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5},
		{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5},
		{3, 1}, {3, 2}, {3, 3}, {3, 4}, {3, 5},
		{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5},
		{5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5},
		{6, 0}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5},
		{7, 1}, {7, 2}, {7, 3}, {7, 4}, {7, 5},
		{8, 0}, {8, 1}, {8, 2}, {8, 3}, {8, 4}, {8, 5},

		//{3, 1}, {3, 2}, {4, 3}, {5, 1}, {5, 2}, {4, 1}, {4, 2},
	}

	for _, p := range table {
		h.Make(p[1], p[0])
	}

	h.Draw()

	js.Global().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		e := args[0]
		x := e.Get("offsetX").Int()
		y := e.Get("offsetY").Int()

		h.Test(x, y)
		return nil
	}))

	stage := factoryBrowser.NewStage()
	stage.Append(h.svg)

	done := make(chan struct{})
	done <- struct{}{}
}
