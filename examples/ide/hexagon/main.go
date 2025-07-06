package main

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/factoryFontFamily"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConversion"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/hexagon"
	"github.com/helmutkemper/webassembly/utilsDraw"
	"github.com/helmutkemper/webassembly/utilsText"
	"github.com/helmutkemper/webassembly/utilsWindow"
	"math"
	"strings"
	"syscall/js"
	"time"
)

type CalcSystem interface {
	GetColRow() (col, row int)
	GetLayout() (layout hexagon.Layout)
	GetCenter() (x, y int)
	SetRowCol(col, row int)
	GetPath() (path []string)
	GetPoints() (points [][2]int)
}

type DrawCell interface {
	Init()
	SetColRow(col, row int)
	SetImg(src string, width, height rulesDensity.Density)
	SetText(text string)
	GetElement() (tagCanvas any)
	ToPng() (img js.Value)
}

type CanvasCell struct {
	canvas                    *html.TagCanvas
	canvasWidth, canvasHeight int

	fontIcon   string
	fontFamily string
	fontSize   int
	fontWeight html.FontWeightRule
	fontStyle  html.FontStyleRule

	calcSystem CalcSystem
}

func (e *CanvasCell) SetWidth(width int) {
	e.canvasWidth = width
}

func (e *CanvasCell) SetHeight(height int) {
	e.canvasHeight = height
}

func (e *CanvasCell) SetCalcSystem(calcSystem CalcSystem) {
	e.calcSystem = calcSystem
}

func (e *CanvasCell) CanvasInit() {
	e.fontSize = 24
	e.fontIcon = utilsText.KFontAwesomeSolid
	e.fontFamily = factoryFontFamily.NewArial()
	e.fontWeight = html.KFontWeightRuleNormal
	e.fontStyle = html.KFontStyleRuleNormal

	e.canvas = factoryBrowser.NewTagCanvas(e.canvasWidth, e.canvasHeight).
		Import("canvas").
		StrokeStyle("red") // todo: tirar daqui

	e.canvas.Font(
		html.Font{
			Style:  e.fontStyle,
			Weight: e.fontWeight,
			Size:   e.fontSize,
			Family: e.fontFamily,
		},
	)
}

func (e *CanvasCell) Init() {

}

// SetColRow
//
// Sets the column and row in the calculation system and updates the Canvas path.
//
//	Note:
//	  * This function is called for each hexagon
func (e *CanvasCell) SetColRow(col, row int) {
	e.calcSystem.SetRowCol(col, row)
	points := e.calcSystem.GetPoints()

	e.canvas.BeginPath()
	for k, point := range points {

		if k == 0 {
			e.canvas.MoveTo(point[0], point[1])
			continue
		}

		e.canvas.LineTo(point[0], point[1])
	}

	e.canvas.LineTo(points[0][0], points[0][1])
	e.canvas.Stroke()
}

func (e *CanvasCell) SetText(text string) {
	fontWeight := e.fontWeight.String()
	fontStyle := e.fontStyle.String()

	width, height := utilsText.GetTextSize(text, e.fontFamily, fontWeight, fontStyle, e.fontSize)

	cx, cy := e.calcSystem.GetCenter()

	x := cx - width/2
	y := cy + height/2 - height/5
	e.canvas.FillStyle("blue")
	e.canvas.FillText(text, x, y, 0)
}

func (e *CanvasCell) GetElement() (tagCanvas any) {
	return e.canvas
}

type SvgCell struct {
	svg      *html.TagSvg
	svgGroup *html.TagSvgG
	svgPath1 *html.TagSvgPath
	svgPath2 *html.TagSvgPath
	svgText  *html.TagSvgText
	svgImg   *html.TagSvgImage

	iconBorder rulesDensity.Density
	iconBottom rulesDensity.Density
	cellHeight rulesDensity.Density

	fontIcon   string
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
	e.iconBorder = rulesDensity.Convert(5)

	e.cellHeight = rulesDensity.Convert(int(math.Sqrt(3) * e.calcSystem.GetLayout().Size.X))

	e.fontSize = 16
	e.fontFamily = factoryFontFamily.NewArial()
	e.fontIcon = utilsText.KFontAwesomeSolid
	e.fontWeight = html.KFontWeightRuleNormal
	e.fontStyle = html.KFontStyleRuleNormal

	e.svg = factoryBrowser.NewTagSvg().
		Append(
			factoryBrowser.NewTagSvgFilter().Id("blur1").Append(
				factoryBrowser.NewTagSvgFeOffset().Dx(1).Dy(1),
				factoryBrowser.NewTagSvgFeBlend().In2(html.KSvgIn2SourceAlpha),
				factoryBrowser.NewTagSvgFeGaussianBlur().
					StdDeviation(3).In(html.KSvgInStrokePaint),
			),
			//factoryBrowser.NewTagSvgFilter().Id("blur2").Append(
			//	factoryBrowser.NewTagSvgFeOffset().Dx(2).Dy(2),
			//	factoryBrowser.NewTagSvgFeBlend().In2("offOut"),
			//	factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(2),
			//),
			//factoryBrowser.NewTagSvgFilter().Id("blur3").Append(
			//	factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(0.5),
			//),
		)

	e.svgGroup = factoryBrowser.NewTagSvgG()

	e.svgPath1 = factoryBrowser.NewTagSvgPath().
		Fill("black").
		Stroke("black").
		StrokeWidth(1).
		Filter("url(#blur1)")
	e.svgGroup.Append(e.svgPath1)

	e.svgPath2 = factoryBrowser.NewTagSvgPath().
		Fill("yellow").
		Stroke("none").
		StrokeWidth(1)
	e.svgGroup.Append(e.svgPath2)

	e.svgImg = factoryBrowser.NewTagSvgImage()
	e.svgGroup.Append(e.svgImg)

	e.svgText = factoryBrowser.NewTagSvgText().
		FontSize(e.fontSize).
		FontFamily(e.fontFamily).
		FontWeight(e.fontWeight).
		FontStyle(e.fontStyle)
	e.svgGroup.Append(e.svgText)

	e.svg.Append(e.svgGroup)
}

func (e *SvgCell) ToPng() (img js.Value) {
	return e.svg.ToPng()
}

// SetColRow
//
// Sets the column and row in the calculation system and updates the SVG path.
//
//	Note:
//	  * This function is called for each hexagon
func (e *SvgCell) SetColRow(col, row int) {
	e.calcSystem.SetRowCol(col, row)
	path := e.calcSystem.GetPath()
	e.svgPath1.D(path)
	e.svgPath2.D(path)
}

func (e *SvgCell) SetImg(src string, width, height rulesDensity.Density) {
	x, y := e.calcSystem.GetCenter()

	y = y - int(e.cellHeight/2) + e.iconBorder.GetInt()

	e.svgImg.HRef(src).
		X(x - width.GetInt()/2).
		Y(y).
		Width(width.GetInt()).
		Height(height.GetInt())

	e.iconBottom = rulesDensity.Convert(int(height.GetFloat() + float64(y)))
}

func (e *SvgCell) SetText(text string) {
	fontWeight := e.fontWeight.String()
	fontStyle := e.fontStyle.String()

	cx, cy := e.calcSystem.GetCenter()

	hexBottomWidth := e.calcSystem.GetLayout().Size.X
	textHorizontalSpace := rulesDensity.Convert(15)
	iconSpace := rulesDensity.Convert(100) + 2*e.iconBorder
	yHexagonTop := rulesDensity.Convert(cy) - e.cellHeight/2
	yHexagonBottom := float64(rulesDensity.Convert(cy) + e.cellHeight/2)

	var calcWidth float64
	var textHeight int
	var yText rulesDensity.Density
	var counterLine = 0
	for {
		_, textHeight = utilsText.GetTextSize(text, e.fontFamily, fontWeight, fontStyle, e.fontSize)

		// TEXT INSIDE HEX
		// ____    +--------- yHexBottomToTextBottom
		// \   |   |  |   / <- triangleSpace
		//  \  | <-+  |  /
		//   \ |      | /
		//    \|______|/

		yText = yHexagonTop + iconSpace + rulesDensity.Convert(textHeight*counterLine) + textHorizontalSpace
		yHexBottomToTextBottom := yHexagonBottom - yText.GetFloat()
		triangleSpace := math.Sin(math.Pi/5) * yHexBottomToTextBottom

		calcWidth = 2*triangleSpace + hexBottomWidth - 5 //todo: 5 = border
		if calcWidth <= hexBottomWidth {
			return
		}

		lines := utilsText.WrapText(text, calcWidth, e.fontFamily, fontWeight, fontStyle, e.fontSize)
		if len(lines) == 0 {
			return
		}

		counterLine += 1
		line := lines[0]
		text = strings.Replace(text, line, "", 1)

		//svgLine := factoryBrowser.NewTagSvgLine().
		//	StrokeWidth(2).
		//	Stroke("red").
		//	X1(cx - int(calcWidth/2.0)).
		//	X2(cx + int(calcWidth/2.0)).
		//	Y1(yText.GetInt()).
		//	Y2(yText.GetInt())
		//e.svgGroup.Append(svgLine)

		textWidth, _ := utilsText.GetTextSize(line, e.fontFamily, fontWeight, fontStyle, e.fontSize)

		svgText := factoryBrowser.NewTagSvgText().
			FontSize(e.fontSize).
			FontFamily(e.fontFamily).
			FontStyle(fontStyle).
			FontWeight(fontWeight).
			X(cx - textWidth/2).
			Y(yText.GetInt()).
			Text(line)

		e.svgGroup.Append(svgText)
	}
}

func (e *SvgCell) GetElement() (tagSvg any) {
	return e.svg
}

// Hexagon represents a hexagonal grid cell in both pixel and coordinate space.
// It captures various properties such as position, layout, and cube coordinates.
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

// Init initializes the hexagon with its layout based on the specified origin (x, y) and size.
func (e *Hexagon) Init(x, y, size int) {
	e.layout = hexagon.Layout{
		Orientation: hexagon.LayoutFlat,
		Size:        hexagon.Point{X: float64(size), Y: float64(size)},
		Origin:      hexagon.Point{X: float64(x), Y: float64(y)},
	}
}

func (e *Hexagon) GetLayout() (layout hexagon.Layout) {
	return e.layout
}

// GetColRow returns the column and row indices of the hexagon in the grid coordinate system.
func (e *Hexagon) GetColRow() (col, row int) {
	return e.col, e.row
}

// GetCenter returns the pixel coordinates (x, y) of the center of the hexagon.
func (e *Hexagon) GetCenter() (x, y int) {
	return e.cx, e.cy
}

// AdjustCenter recalculates and returns the adjusted pixel coordinates (cx, cy) for the center of a hexagon based on inputs (x, y).
func (e *Hexagon) AdjustCenter(x, y int) (cx, cy int) {
	hex := e.colHexToRow(hexagon.Point{X: float64(x), Y: float64(y)})
	point := hexagon.HexToPixel(e.layout, hex)
	return rulesConversion.FloatToInt(point.X), rulesConversion.FloatToInt(point.Y)
}

// SetPixelXY sets the hexagon's column and row based on the provided pixel coordinates (x, y).
func (e *Hexagon) SetPixelXY(x, y int) {
	hex := e.colHexToRow(hexagon.Point{X: float64(x), Y: float64(y)})
	cord := hexagon.QDoubledFromCube(hex)
	e.SetRowCol(cord.Col, cord.Row)
}

// SetRowCol updates the hexagon's column and row indices and triggers conversion of coordinates and layout adjustments.
func (e *Hexagon) SetRowCol(col, row int) {
	e.col = col
	e.row = row
	e.convertManager(e.col, e.row)
}

// colHexToRow converts a 2D pixel position (Point) to a hexagonal grid coordinate (Hex) based on the instance layout.
func (e *Hexagon) colHexToRow(point hexagon.Point) (hex hexagon.Hex) {
	return hexagon.PixelToHex(e.layout, point)
}

// colRowToHex converts a column and row from a doubled coordinate system to a cube coordinate (hexagon.Hex).
func (e *Hexagon) colRowToHex(col, row int) (hex hexagon.Hex) {
	return hexagon.QDoubledToCube(hexagon.DoubledCoordinate{Col: col, Row: row})
}

// convertManager recalculates and updates hexagon's attributes based on provided column and row indices in the grid system.
func (e *Hexagon) convertManager(col, row int) {
	e.col = col
	e.row = row
	e.hex = e.colRowToHex(col, row)
	e.point = hexagon.HexToPixel(e.layout, e.hex)
	e.cx = rulesConversion.FloatToInt(e.point.X)
	e.cy = rulesConversion.FloatToInt(e.point.Y)
}

// GetPath generates a path for the hexagon's outline as a series of SVG-compatible commands based on its corners.
func (e *Hexagon) GetPath() (path []string) {
	//points := hexagon.PolygonCorners(e.layout, e.hex)

	points := utilsDraw.Polygon(6, rulesDensity.Convert(int(e.layout.Size.X)-5), rulesDensity.Convert(e.cx), rulesDensity.Convert(e.cy), 0)

	for k, point := range points {
		if k == 0 {
			path = append(path, fmt.Sprintf("M %.2f,%.2f ", point[0], point[1]))
			continue
		}

		path = append(path, fmt.Sprintf("L %.2f,%.2f ", point[0], point[1]))
	}
	path = append(path, "z")
	return
}

// GetPoints returns the 2D integer coordinates of the hexagon's corners based on its layout and position.
func (e *Hexagon) GetPoints() (points [][2]int) {
	ps := hexagon.PolygonCorners(e.layout, e.hex)
	points = make([][2]int, len(ps))
	for k, point := range ps {
		points[k] = [2]int{int(point.X), int(point.Y)}
	}

	return
}

type HexagonDraw struct {
	svg    *html.TagSvg
	canvas *html.TagCanvas
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

	e.svg = factoryBrowser.NewTagSvg()
	e.canvas = new(html.TagCanvas)
}

func (e *HexagonDraw) DrawText(text string) {
	e.drawSystem.SetText(text)
}

func (e *HexagonDraw) Draw(col, row int) {
	e.drawSystem.Init()
	e.drawSystem.SetColRow(col, row)

	if converted, ok := e.drawSystem.GetElement().(*html.TagSvg); ok {
		e.svg.Append(converted)
	}
}

func (e *HexagonDraw) SetImg(src string, width, height rulesDensity.Density) {
	e.drawSystem.SetImg(src, width, height)
}

func (e *HexagonDraw) GetSvg() (tagSvg *html.TagSvg) {
	return e.svg
}

type CellMenu struct {
	Col        int
	Row        int
	Icon       string
	IconWidth  rulesDensity.Density
	IconHeight rulesDensity.Density
	Text       string
}

func main() {

	utilsWindow.InjectBodyNoMargin()
	utilsText.InjectFontAwesomeCSS()

	time.Sleep(100 * time.Millisecond)

	//document := js.Global().Get("document")

	screenWidth, screenHeight := utilsWindow.GetScreenSize()

	stage := factoryBrowser.NewStage()

	size := 100
	hex := new(Hexagon)
	hex.Init(0, 0, size)

	cellSvg := new(SvgCell)
	cellSvg.SetCalcSystem(hex)

	cellCanvas := new(CanvasCell)
	cellCanvas.SetCalcSystem(hex)
	cellCanvas.SetWidth(screenWidth)
	cellCanvas.SetHeight(screenHeight)
	cellCanvas.CanvasInit()

	hexSvg := new(HexagonDraw)
	hexSvg.SetDrawSystem(cellSvg)
	hexSvg.Init()

	//hexCanvas := new(HexagonDraw)
	//hexCanvas.SetDrawSystem(cellCanvas)
	//hexCanvas.Init()

	mainSvg := factoryBrowser.NewTagSvg().Import("svg").ResizeToWindow()

	//cellList := []CellMenu{
	//	{Col: 2, Row: 2, Icon: "./icon.png", IconWidth: 100, IconHeight: 100, Text: "simple loop"},
	//	//{Col: 2, Row: 4, Icon: "./icon.png", IconWidth: 100, IconHeight: 100, Text: "loop"},
	//	//{Col: 1, Row: 3, Icon: "./icon.png", IconWidth: 100, IconHeight: 100, Text: "loop"},
	//	//{Col: 3, Row: 3, Icon: "./icon.png", IconWidth: 100, IconHeight: 100, Text: "loop"},
	//	//{1, 3}, {3, 3},
	//	//{2, 4},
	//	//{1, 5}, {3, 5},
	//	//{2, 6},
	//}

	//for _, cell := range cellList {
	//	hexSvg.Draw(cell.Col, cell.Row)
	//
	//	hexSvg.SetImg(cell.Icon, cell.IconWidth, cell.IconHeight)
	//	hexSvg.DrawText(cell.Text)
	//	mainSvg.Append(hexSvg.GetSvg())
	//
	//	//hexCanvas.Draw(col, row)
	//	//hexCanvas.DrawText(fmt.Sprintf("%v, %v", col, row))
	//}

	for col := 0; col < int(float64(screenWidth)/(float64(size)*2.0*3.0/4.0))+2; col += 1 {
		for row := 0; row < int(float64(screenHeight)/(float64(size)*math.Sqrt(3))+2)*2; row += 1 {

			if (col+row)%2 != 0 {
				continue
			}

			hexSvg.Draw(col, row)
			hexSvg.DrawText(fmt.Sprintf("%v,%v", col, row))
			mainSvg.Append(hexSvg.GetSvg())

			//hexCanvas.Draw(col, row)
			//hexCanvas.DrawText(fmt.Sprintf("%v, %v", col, row))
			//time.Sleep(time.Nanosecond)
		}
	}

	img := factoryBrowser.NewTagImg().Import("img")
	time.Sleep(1000 * time.Millisecond)
	img.Src(mainSvg.ToPng(), true)
	stage.Append(img)

	stage.Append(mainSvg)
	//url := mainSvg.ToPngResized(0.5, 0.5)

	//document.Call("getElementById", "test").Set("src", url)
	//document.Call("getElementById", "test").Get("style").Set("width", mainSvg.GetWidth())
	//document.Call("getElementById", "test").Get("style").Set("height", mainSvg.GetHeight())

	//document.Call("getElementById", "test").Set("src", url)
	//document.Call("getElementById", "test").Get("style").Set("width", mainSvg.GetWidth())
	//document.Call("getElementById", "test").Get("style").Set("height", mainSvg.GetHeight())

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
