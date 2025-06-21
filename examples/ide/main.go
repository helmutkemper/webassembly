package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/browser/stage"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/examples/ide/devices/block"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
	"github.com/helmutkemper/webassembly/hexagon"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/utilsText"
	"github.com/helmutkemper/webassembly/utilsWindow"
	"image/color"
	"log"
	"math"
	"syscall/js"
)

type MenuOptions struct {
	Label string `wasmPanel:"type:label"`
	Icon  string `wasmPanel:"type:icon"`
	//IconLeft  string        `wasmPanel:"type:iconLeft"`
	//IconRight string        `wasmPanel:"type:iconRight"`
	Type    string        `wasmPanel:"type:type"`
	Items   []MenuOptions `wasmPanel:"type:options"`
	Action  js.Func       `wasmPanel:"type:action"`
	Submenu []MenuOptions `wasmPanel:"type:subMenu"`
}

type ComponentControlPanel struct {
	components.Components
	components.MainMenu

	Menu *[]MenuOptions `wasmPanel:"type:mainMenu;func:InitMainMenu;label:Menu;top:200;left:5;columns:3"`
}

func (e *ComponentControlPanel) InitMainMenu() {
	e.Menu = getMenuComplex()
}

func (e *ComponentControlPanel) Init() (panel *html.TagDiv, err error) {
	panel, err = e.Components.Init(e)
	return
}

func getMenuComplex() (options *[]MenuOptions) {
	return &[]MenuOptions{
		{
			Type:  "label",
			Label: "Statement:",
		},
		{
			Type: "grid",
			Items: []MenuOptions{
				{
					Type:   "grid",
					Label:  "Loop",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
					Submenu: []MenuOptions{
						{
							Type: "grid",
							Items: []MenuOptions{
								{
									Label:  "Basic loop",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
								{
									Label:  "Basic loop",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
								{
									Label:  "Basic loop",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
							},
						},
					},
				},
				{
					Type:   "grid",
					Label:  "Math",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
					Submenu: []MenuOptions{
						{
							Type: "grid",
							Items: []MenuOptions{
								{
									Label:  "Add",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
								{
									Label:  "Sub",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
								{
									Label:  "Mul",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
								{
									Label:  "Div",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
							},
						},
					},
				},
			},
		},
	}
}

var GlobalControlPanel = new(ComponentControlPanel)
var mainStage *stage.Stage

func main() {
	var err error

	utilsWindow.InjectBodyNoMargin()
	utilsText.InjectFontAwesomeCSS()

	//mainStage = factoryBrowser.NewStage()

	screenWidth, screenHeight := utilsWindow.GetScreenSize()
	mainSvg := factoryBrowser.NewTagSvg().
		DataKey("name", "mainSvg").
		Import(rulesStage.KStageId).
		X(0).
		Y(0).
		Width(screenWidth).
		Height(screenHeight).
		PreserveAspectRatio(html.KRatioXMinYMin, html.KMeetOrSliceReferenceMeet).
		AddStyle("display", "block") //.
	//ViewBox([]float64{0, 0, 0.5 * float64(screenWidth), 0.5 * float64(screenHeight)})

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	mainSvg.ViewBox([]float64{0, 0, 0.5 * float64(screenWidth), 0.5 * float64(screenHeight)})
	//}()

	size := 30
	hex := new(rulesStage.Hexagon)
	hex.Init(0, 0, size)

	cellCanvas := new(CanvasCell)
	cellCanvas.SetCalcSystem(hex)
	cellCanvas.SetWidth(screenWidth)
	cellCanvas.SetHeight(screenHeight)
	cellCanvas.CanvasInit()

	hexCanvas := new(HexagonDraw)
	hexCanvas.SetDrawSystem(cellCanvas)
	hexCanvas.Init()

	for col := 0; col < int(float64(screenWidth)/(float64(size)*2.0*3.0/4.0))+2; col += 1 {
		for row := 0; row < int(float64(screenHeight)/(float64(size)*math.Sqrt(3))+2)*2; row += 1 {

			if (col+row)%2 != 0 {
				continue
			}

			hexCanvas.Draw(col, row)
			//hexCanvas.DrawText(fmt.Sprintf("%v, %v", col, row))
			//time.Sleep(time.Nanosecond)
		}
	}

	resizeButton := new(block.ResizeButtonHexagon)
	resizeButton.SetSize(20)
	resizeButton.SetSides(6)
	resizeButton.SetFillColor("red")
	resizeButton.SetStrokeColor("green")
	resizeButton.SetStrokeWidth(2)
	//resizeButton.SetCX(30)
	//resizeButton.SetCY(30)
	//resizeButton.Init()

	//mainStage.Append(resizeButton.GetSvg())

	stmLoop := new(devices.StatementLoop)
	stmLoop.SetResizeButton(resizeButton)
	stmLoop.SetGridAdjust(hex)
	stmLoop.SetMainSvg(mainSvg)
	stmLoop.SetPosition(50, 50)
	_ = stmLoop.Init()
	//url := stmLoop.ToPng()

	//stmAdd := new(devices.StatementAdd)
	//stmAdd.SetPosition(300, 300)
	//_ = stmAdd.Init()
	//url := stmAdd.ToPng()

	if _, err = GlobalControlPanel.Init(); err != nil {
		panic(err)
	}

	//document := js.Global().Get("document")
	//document.Call("getElementById", "test").Set("src", url)
	//document.Call("getElementById", "test").Get("style").Set("width", stmLoop.GetWidth())
	//document.Call("getElementById", "test").Get("style").Set("height", stmLoop.GetHeight())

	//mainStage.Append(panel)

	factoryBrowser.NewTagImg().Import("imgTest").Src(mainSvg.ToPng(), true)

	doc := js.Global().Get("document")

	canvas := doc.Call("getElementById", "canvas")
	canvas.Get("style").Set("position", "absolute")
	canvas.Get("style").Set("top", "0")
	canvas.Get("style").Set("left", "0")
	canvas.Get("style").Set("zIndex", "0")

	svg := doc.Call("getElementById", "hexagonIde")
	svg.Get("style").Set("position", "absolute")
	svg.Get("style").Set("top", "0")
	svg.Get("style").Set("left", "0")
	svg.Get("style").Set("zIndex", "1")
	svg.Get("style").Set("pointerEvents", "auto")

	img := doc.Call("getElementById", "imgTest")
	img.Get("style").Set("position", "absolute")
	img.Get("style").Set("top", "0")
	img.Get("style").Set("left", "0")
	img.Get("style").Set("zIndex", "-1")
	img.Get("style").Set("visibility", "hidden")

	done := make(chan struct{})
	<-done
}

type CanvasCell struct {
	canvas                    *html.TagCanvas
	canvasWidth, canvasHeight int

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
	e.fontFamily = utilsText.KFontAwesomeSolid
	e.fontWeight = html.KFontWeightRuleNormal
	e.fontStyle = html.KFontStyleRuleNormal

	e.canvas = factoryBrowser.NewTagCanvas(e.canvasWidth, e.canvasHeight).
		Import("canvas").
		StrokeStyle(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x20}) // todo: tirar daqui

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

type DrawCell interface {
	Init()
	SetColRow(col, row int)
	SetText(text string)
	GetElement() (tagCanvas any)
}

type CalcSystem interface {
	GetColRow() (col, row int)
	GetCenter() (x, y int)
	SetRowCol(col, row int)
	GetPath() (path []string)
	GetPoints() (points [][2]int)
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

func (e *HexagonDraw) GetSvg() (tagSvg *html.TagSvg) {
	return e.svg
}
