package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/browser/stage"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/examples/ide/devices/block"
	"github.com/helmutkemper/webassembly/examples/ide/hexagonMenu"
	"github.com/helmutkemper/webassembly/examples/ide/manager"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
	"github.com/helmutkemper/webassembly/examples/ide/splashScreen"
	"github.com/helmutkemper/webassembly/hexagon"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/utilsText"
	"github.com/helmutkemper/webassembly/utilsWindow"
	"image/color"
	"log"
	"math"
	"syscall/js"
	"time"
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
		Width(rulesDensity.Density(screenWidth).GetInt()).
		Height(rulesDensity.Density(screenHeight).GetInt())

	splash := new(splashScreen.Control)
	splash.Init(mainSvg)

	splash.AddText("Creating Density System for Multiple Screens")

	size := rulesDensity.Density(3)
	hex := new(rulesStage.Hexagon)
	hex.Init(0, 0, size)

	splash.AddText("Creating Coordinate System")

	//cellCanvas := new(CanvasCell)
	//cellCanvas.SetCalcSystem(hex)
	//cellCanvas.SetWidth(rulesDensity.Density(screenWidth))
	//cellCanvas.SetHeight(rulesDensity.Density(screenHeight))
	//cellCanvas.CanvasInit()

	splash.AddText("Drawing work environment")

	//hexCanvas := new(HexagonDraw)
	//hexCanvas.SetDrawSystem(cellCanvas)
	//hexCanvas.Init()

	colT := int(rulesDensity.Density(screenWidth).GetFloat()/(float64(size)*2.0*3.0/4.0)) + 2
	rowT := int(rulesDensity.Density(screenHeight).GetFloat()/(float64(size)*math.Sqrt(3))+2) * 2

	counter := 0
	for col := 0; col < colT; col += 1 {
		for row := 0; row < rowT; row += 1 {
			switch float64(counter) {
			case float64(colT*rowT) * 0.2:
				splash.AddText("Drawing work environment 20%")
				time.Sleep(time.Nanosecond)
			case float64(colT*rowT) * 0.4:
				splash.AddText("Drawing work environment 40%")
				time.Sleep(time.Nanosecond)
			case float64(colT*rowT) * 0.6:
				splash.AddText("Drawing work environment 60%")
				time.Sleep(time.Nanosecond)
			case float64(colT*rowT) * 0.8:
				splash.AddText("Drawing work environment 80%")
				time.Sleep(time.Nanosecond)
			}
			counter += 1

			if (col+row)%2 != 0 {
				continue
			}

			//hexCanvas.Draw(col, row)
			//hexCanvas.DrawText(fmt.Sprintf("%v, %v", col, row))
			//time.Sleep(time.Nanosecond)
		}
	}
	splash.AddText("Drawing work environment 100%")

	splash.AddText("Preparing design elements")

	resizeButton := new(block.ResizeButtonHexagon)
	resizeButton.SetSize(10)
	resizeButton.SetSpace(30)
	resizeButton.SetSides(6)
	resizeButton.SetFillColor("red")
	resizeButton.SetStrokeColor("green")
	resizeButton.SetStrokeWidth(2)
	//resizeButton.SetRotation(math.Pi / 4)
	//resizeButton.SetCX(30)
	//resizeButton.SetCY(30)
	//resizeButton.Init()

	draggerButton := new(block.ResizeButtonHexagon)
	draggerButton.SetSize(20)
	draggerButton.SetSpace(10)
	draggerButton.SetSides(3)
	draggerButton.SetFillColor(color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0x20})
	draggerButton.SetStrokeColor("none")
	draggerButton.SetStrokeWidth(2)

	//mainStage.Append(resizeButton.GetSvg())

	splash.AddText("Registering devices and functions")

	stmLoop := new(devices.StatementLoop)
	stmLoop.SetResizerButton(resizeButton)
	stmLoop.SetDraggerButton(draggerButton)
	stmLoop.SetGridAdjust(hex)
	stmLoop.SetMainSvg(mainSvg)
	_ = stmLoop.Init()
	manager.Manager.RegisterIcon(stmLoop.GetIcon())

	stmAdd := new(devices.StatementAdd)
	stmAdd.SetResizerButton(resizeButton)
	stmAdd.SetDraggerButton(draggerButton)
	stmAdd.SetGridAdjust(hex)
	stmAdd.SetMainSvg(mainSvg)
	_ = stmAdd.Init()
	//manager.Manager.Register(stmAdd)

	stmSub := new(devices.StatementSub)
	stmSub.SetResizerButton(resizeButton)
	stmSub.SetDraggerButton(draggerButton)
	stmSub.SetGridAdjust(hex)
	stmSub.SetMainSvg(mainSvg)
	_ = stmSub.Init()
	//manager.Manager.Register(stmSub)

	stmDiv := new(devices.StatementDiv)
	stmDiv.SetResizerButton(resizeButton)
	stmDiv.SetDraggerButton(draggerButton)
	stmDiv.SetGridAdjust(hex)
	stmDiv.SetMainSvg(mainSvg)
	_ = stmDiv.Init()
	//manager.Manager.Register(stmDiv)

	stmMul := new(devices.StatementMul)
	stmMul.SetResizerButton(resizeButton)
	stmMul.SetDraggerButton(draggerButton)
	stmMul.SetGridAdjust(hex)
	stmMul.SetMainSvg(mainSvg)
	_ = stmMul.Init()
	//manager.Manager.Register(stmMul)

	splash.Hide()

	//-------------------------------------------------
	stmLoop.SetPosition(150, 50)
	//stmLoop.Append()

	stmAdd.SetPosition(300, 150)
	//stmAdd.Append()

	stmSub.SetPosition(400, 250)
	//stmSub.Append()

	stmDiv.SetPosition(400, 450)
	//stmDiv.Append()

	stmMul.SetPosition(200, 450)
	//stmMul.Append()

	if _, err = GlobalControlPanel.Init(); err != nil {
		panic(err)
	}

	//mainStage.Append(panel)

	//ic := new(hexagonMenu.MakeIcon).GoBack()
	//factoryBrowser.NewTagImg().Import("imgTest").Src(ic.GetIcon(false), true)

	hexagonMenu.Menu.Process(mainSvg)

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
	img.Get("style").Set("top", "100")
	img.Get("style").Set("left", "100")
	img.Get("style").Set("zIndex", "10")
	//img.Get("style").Set("visibility", "hidden")

	//for a := 0; a != 300; a += 15 {
	//	for b := 0; b != 300; b += 15 {
	//		var cxD, cyD = rulesDensity.Density(a), rulesDensity.Density(b)
	//		cx, cy := hex.AdjustCenter(cxD.GetInt(), cyD.GetInt())
	//		rec := factoryBrowser.NewTagSvg().
	//			Append(
	//				factoryBrowser.NewTagSvgRect().
	//					X(cx - 3).
	//					Y(cy - 3).
	//					Width(6).
	//					Height(6).
	//					Fill("red"),
	//			)
	//		mainSvg.Append(rec)
	//	}
	//}

	done := make(chan struct{})
	<-done
}

type CanvasCell struct {
	canvas                    *html.TagCanvas
	canvasWidth, canvasHeight rulesDensity.Density

	fontFamily string
	fontSize   rulesDensity.Density
	fontWeight html.FontWeightRule
	fontStyle  html.FontStyleRule

	calcSystem CalcSystem
}

func (e *CanvasCell) SetWidth(width rulesDensity.Density) {
	e.canvasWidth = width
}

func (e *CanvasCell) SetHeight(height rulesDensity.Density) {
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

	e.canvas = factoryBrowser.NewTagCanvas(e.canvasWidth.GetInt(), e.canvasHeight.GetInt()).
		Import("canvas").
		StrokeStyle(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x20}) // todo: tirar daqui

	e.canvas.Font(
		html.Font{
			Style:  e.fontStyle,
			Weight: e.fontWeight,
			Size:   e.fontSize.GetInt(),
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
			e.canvas.MoveTo(point[0].GetInt(), point[1].GetInt())
			continue
		}

		e.canvas.LineTo(point[0].GetInt(), point[1].GetInt())
	}

	e.canvas.LineTo(points[0][0].GetInt(), points[0][1].GetInt())
	e.canvas.Stroke()
}

func (e *CanvasCell) SetText(text string) {
	fontWeight := e.fontWeight.String()
	fontStyle := e.fontStyle.String()

	widthInt, heightInt := utilsText.GetTextSize(text, e.fontFamily, fontWeight, fontStyle, e.fontSize.GetInt())

	cx, cy := e.calcSystem.GetCenter()

	x := cx - rulesDensity.Convert(widthInt)/2
	y := cy + rulesDensity.Convert(heightInt)/2 - rulesDensity.Convert(heightInt)/5
	e.canvas.FillStyle("blue")
	e.canvas.FillText(text, x.GetInt(), y.GetInt(), 0)
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
	GetCenter() (x, y rulesDensity.Density)
	SetRowCol(col, row int)
	GetPath() (path []string)
	GetPoints() (points [][2]rulesDensity.Density)
}

type HexagonDraw struct {
	svg    *html.TagSvg
	canvas *html.TagCanvas
	sides  int
	space  rulesDensity.Density
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

/*
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
<!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.-->
<path d=

"M0 224c0 17.7 14.3 32 32 32s32-14.3 32-32c0-53 43-96 96-96l160 0 0 32c0 12.9 7.8 24.6 19.8 29.6s25.7 2.2 34.9-6.9l64-64c12.5-12.5 12.5-32.8 0-45.3l-64-64c-9.2-9.2-22.9-11.9-34.9-6.9S320 19.1 320 32l0 32L160 64C71.6 64 0 135.6 0 224zm512 64c0-17.7-14.3-32-32-32s-32 14.3-32 32c0 53-43 96-96 96l-160 0 0-32c0-12.9-7.8-24.6-19.8-29.6s-25.7-2.2-34.9 6.9l-64 64c-12.5 12.5-12.5 32.8 0 45.3l64 64c9.2 9.2 22.9 11.9 34.9 6.9s19.8-16.6 19.8-29.6l0-32 160 0c88.4 0 160-71.6 160-160z"

/></svg>
<svg xmlns="http://www.w3.org/2000/svg" x="100" y="100" width="200" height="200"><style xmlns="http://www.w3.org/2000/svg">
@font-face {
	font-family: "FASolid";
	src: url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/webfonts/fa-solid-900.woff2") format("woff2");
	font-style: normal;
}
</style><path stroke-width="4" stroke="rgba(95,95,95,1)" fill="rgba(180,180,255,1)" d="M 200 100 L 150 186.60254 L 50 186.60254 L 0 100 L 50 13.39746 L 150 13.39746 L 200 100 z"></path><text xmlns="http://www.w3.org/2000/svg" font-family="FASolid" font-size="90" fill="white" x="55" y="110"></text><text xmlns="http://www.w3.org/2000/svg" font-family="Helvetica" font-size="20" fill="rgba(0,0,0,1)" x="75" y="160">Loop</text></svg>
*/
