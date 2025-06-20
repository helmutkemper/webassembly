package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/browser/stage"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/textUtil"
	"github.com/helmutkemper/webassembly/windowUtils"
	"log"
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

	windowUtils.InjectBodyNoMargin()
	textUtil.InjectFontAwesomeCSS()

	screenWidth, screenHeight := windowUtils.GetScreenSize()
	mainSvg := factoryBrowser.NewTagSvg().
		Import(rulesStage.KStageId).
		X(0).
		Y(0).
		Width(screenWidth).
		Height(screenHeight)

	grid := new(rulesStage.Hexagon)
	grid.Init(0, 0, 40)

	stmLoop := new(devices.StatementLoop)
	stmLoop.SetGridAdjust(grid)
	stmLoop.SetPosition(50, 450)
	_ = stmLoop.Init()
	//url := stmLoop.ToPng()
	stmLoop.SetWarning(true)

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

	done := make(chan struct{})
	<-done
}
