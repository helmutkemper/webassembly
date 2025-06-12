package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/platform/components"
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

func main() {

	var err error
	var panel *html.TagDiv

	stmLoop := new(devices.StatementLoop)
	stmLoop.SetPosition(50, 50)
	_ = stmLoop.Init()
	//stmLoop.SetWarning(true)

	stmAdd := new(devices.StatementAdd)
	stmAdd.SetPosition(200, 200)
	_ = stmAdd.Init()

	stage := factoryBrowser.NewStage()
	stage.Append(stmLoop.Get())
	stage.Append(stmAdd.Get())

	if panel, err = GlobalControlPanel.Init(); err != nil {
		panic(err)
	}

	stage.Append(panel)

	done := make(chan struct{})
	<-done
}
