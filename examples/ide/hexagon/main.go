package main

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConversion"
	"github.com/helmutkemper/webassembly/hexagon"
	"github.com/helmutkemper/webassembly/textUtil"
	"log"
)

type CalcSystem interface {
	GetColRow() (col, row int)
	GetCenter() (x, y int)
	SetRowCol(col, row int)
	GetPath() (path []string)
}

type DrawCell interface {
	Init()
	SetColRow(col, row int)
	SetText(text string)
	GetSvg() (tagSvg *html.TagSvg)
}

type SvgCell struct {
	svg      *html.TagSvg
	svgGroup *html.TagSvgG
	svgPath  *html.TagSvgPath
	svgText  *html.TagSvgText

	fontFamily string
	fontSize   int
	fontWeight html.FontWeightRule
	fontStyle  html.FontStyleRule

	calcSystem CalcSystem
}

func (e *SvgCell) SetCalcSystem(calcSystem CalcSystem) {
	e.calcSystem = calcSystem
}

func (e *SvgCell) Init() {
	e.fontSize = 20
	e.fontFamily = textUtil.KFontAwesomeSolid
	e.fontWeight = html.KFontWeightRuleNormal
	e.fontStyle = html.KFontStyleRuleNormal

	e.svg = factoryBrowser.NewTagSvg().
		Append(
			factoryBrowser.NewTagSvgFilter().Id("blur1").Append(
				factoryBrowser.NewTagSvgFeOffset().Dx(1).Dy(1),
				factoryBrowser.NewTagSvgFeBlend().In2(html.KSvgIn2SourceAlpha),
				factoryBrowser.NewTagSvgFeGaussianBlur().
					StdDeviation(2).In(html.KSvgInStrokePaint),
			),
			//	factoryBrowser.NewTagSvgFilter().Id("blur2").Append(
			//		factoryBrowser.NewTagSvgFeOffset().Dx(2).Dy(2),
			//		factoryBrowser.NewTagSvgFeBlend().In2("offOut"),
			//		factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(2),
			//	),
			//	factoryBrowser.NewTagSvgFilter().Id("blur3").Append(
			//		factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(0.5),
			//	),
		)

	e.svgGroup = factoryBrowser.NewTagSvgG()

	e.svgPath = factoryBrowser.NewTagSvgPath().
		Fill("yellow").
		Stroke("black").
		StrokeWidth(1) //.
	//Filter("url(#blur1)")
	e.svgGroup.Append(e.svgPath)

	e.svgText = factoryBrowser.NewTagSvgText().
		FontSize(e.fontSize).
		FontFamily(e.fontFamily).
		FontWeight(e.fontWeight).
		FontStyle(e.fontStyle)
	e.svgGroup.Append(e.svgText)

	e.svg.Append(e.svgGroup)
}

func (e *SvgCell) SetColRow(col, row int) {
	e.calcSystem.SetRowCol(col, row)
	path := e.calcSystem.GetPath()
	e.svgPath.D(path)
}

func (e *SvgCell) SetText(text string) {
	fontWeight := e.fontWeight.String()
	fontStyle := e.fontStyle.String()

	width, height := textUtil.GetTextSize(text, e.fontFamily, fontWeight, fontStyle, e.fontSize)
	log.Printf("w: %v, h: %v", width, height)

	cx, cy := e.calcSystem.GetCenter()

	e.svgText.X(cx - width/2).
		Y(cy + height/2 - height/5).
		Text(text).
		FontSize(e.fontSize).
		FontFamily(e.fontFamily).
		FontWeight(fontWeight)

	//icon := factoryBrowser.NewTagSvgText().
	//	Text(text).
	//	FontSize(fontSize).
	//	FontFamily(e.fontFamily).
	//	FontWeight(fontWeight).
	//	X(cx - width/2).
	//	Y(cy + height/2 - height/5)
	//
	//e.svg.Append(icon)
}

func (e *SvgCell) GetSvg() (tagSvg *html.TagSvg) {
	return e.svg
}

type Hexagon struct {
	// Column and row in the doubled coordinate system
	col, row int

	// X and Y coordinates in 2D space
	cx, cy int

	// Layout defines how hexagons are positioned and transformed on the screen.
	//
	//   - Orientation specifies whether the hexes are pointy-topped or flat-topped,
	//     along with the transformation matrices for converting between hex and pixel space.
	//   - Size defines the radius (width and height) of a single hexagon in screen units (pixels).
	//   - Origin specifies the pixel coordinate where the hex grid origin (0,0,0) will be drawn.
	layout hexagon.Layout

	// Point represents a 2D position in pixel or screen space.
	//
	// It is typically used to store the result of hex-to-pixel conversion,
	// define layout sizes, or calculate hex corner positions for rendering.
	//
	//   - X: horizontal position.
	//   - Y: vertical position.
	point hexagon.Point

	// Hex represents a cube coordinate (q, r, s) used to model positions in a hexagonal grid.
	//
	// Cube coordinates satisfy the constraint q + r + s = 0, allowing for unambiguous positioning
	// and efficient computation of directions, distances, and neighbors.
	//
	//   - Q: corresponds to the "column" axis.
	//   - R: corresponds to the "row" axis.
	//   - S: the third axis, derived from Q and R (S = -Q - R).
	hex hexagon.Hex
}

func (e *Hexagon) Init(x, y, size int) {
	e.layout = hexagon.Layout{
		Orientation: hexagon.LayoutFlat,
		Size:        hexagon.Point{X: float64(size), Y: float64(size)},
		Origin:      hexagon.Point{X: float64(x), Y: float64(y)},
	}
}

func (e *Hexagon) GetColRow() (col, row int) {
	return e.col, e.row
}

func (e *Hexagon) GetCenter() (x, y int) {
	return e.cx, e.cy
}

func (e *Hexagon) SetRowCol(col, row int) {
	e.col = col
	e.row = row
	e.convertManager(e.col, e.row)
}

func (e *Hexagon) colRowToHex(col, row int) (hex hexagon.Hex) {
	return hexagon.QDoubledToCube(hexagon.DoubledCoordinate{Col: col, Row: row})
}

func (e *Hexagon) convertManager(col, row int) {
	e.col = col
	e.row = row
	e.hex = e.colRowToHex(col, row)
	e.point = hexagon.HexToPixel(e.layout, e.hex)
	e.cx = rulesConversion.FloatToInt(e.point.X)
	e.cy = rulesConversion.FloatToInt(e.point.Y)
}

func (e *Hexagon) GetPath() (path []string) {
	points := hexagon.PolygonCorners(e.layout, e.hex)
	for k, point := range points {
		if k == 0 {
			path = append(path, fmt.Sprintf("M %.2f,%.2f ", point.X, point.Y))
			continue
		}

		path = append(path, fmt.Sprintf("L %.2f,%.2f ", point.X, point.Y))
	}
	path = append(path, "z")
	return
}

type HexagonDraw struct {
	svg    *html.TagSvg
	sides  int
	space  int
	radius int
	layout hexagon.Layout

	drawSystem DrawCell
}

func (e *HexagonDraw) SetDrawSystem(system DrawCell) {
	e.drawSystem = system
}

func (e *HexagonDraw) Init() {
	e.sides = 6
	e.space = 10
	e.radius = 100

	e.svg = factoryBrowser.NewTagSvg().
		Width("100vw").
		Height("100vh")
}

func (e *HexagonDraw) DrawText(text string) {
	e.drawSystem.SetText(text)
}

func (e *HexagonDraw) Draw(col, row int) {
	e.drawSystem.Init()
	e.drawSystem.SetColRow(col, row)

	e.svg.Append(e.drawSystem.GetSvg())
}

func (e *HexagonDraw) GetSvg() (tagSvg *html.TagSvg) {
	return e.svg
}

const faStar = "\uf005"

func main() {

	textUtil.InjectFontAwesomeCSS()

	stage := factoryBrowser.NewStage()

	hex := new(Hexagon)
	hex.Init(50, 50, 30)

	cell := new(SvgCell)
	cell.SetCalcSystem(hex)
	cell.Init()

	h := new(HexagonDraw)
	h.SetDrawSystem(cell)
	h.Init()

	table := [][]int{

		{0, 0}, {0, 2}, {0, 4}, {0, 6},
		{1, 1}, {1, 3}, {1, 5},
		{2, 0}, {2, 2}, {2, 4}, {2, 6},
		{3, 1}, {3, 3}, {3, 5},
		{4, 0}, {4, 2}, {4, 4}, {4, 6},
		//{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5},
		//{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5},
		//{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5},
		//{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}, {3, 5},
		//{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5},
		//{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5},
		//{6, 0}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5},
		//{7, 0}, {7, 1}, {7, 2}, {7, 3}, {7, 4}, {7, 5},
		//{8, 0}, {8, 1}, {8, 2}, {8, 3}, {8, 4}, {8, 5},

		//{3, 1}, {3, 2}, {4, 3}, {5, 1}, {5, 2}, {4, 1}, {4, 2},
	}

	for _, p := range table {
		h.Draw(p[1], p[0])
		//h.DrawText(fmt.Sprintf("%v,%v", p[0], p[1]))
		h.DrawText("\uf197")
		stage.Append(h.GetSvg())
	}

	//js.Global().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
	//	e := args[0]
	//	x := e.Get("offsetX").Int()
	//	y := e.Get("offsetY").Int()
	//
	//	h.Test(x, y)
	//	return nil
	//}))

	done := make(chan struct{})
	done <- struct{}{}
}
