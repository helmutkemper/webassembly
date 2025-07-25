package devices

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/connection"
	"github.com/helmutkemper/webassembly/examples/ide/devices/block"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/ornament/math"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/examples/ide/rulesSequentialId"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
	"github.com/helmutkemper/webassembly/platform/components"
	"log"
	"syscall/js"
)

type StatementAdd struct {
	block block.Block
	menu  Menu

	defaultWidth          rulesDensity.Density
	defaultHeight         rulesDensity.Density
	horizontalMinimumSize rulesDensity.Density
	verticalMinimumSize   rulesDensity.Density

	ornamentDraw     *math.OrnamentAdd
	ornamentDrawIcon *math.OrnamentAdd

	id string
	//inputA                *connection.Connection
	//inputB                *connection.Connection
	//output                *connection.Connection
	debugMode bool

	gridAdjust rulesStage.GridAdjust
}

func (e *StatementAdd) Append() {
	e.block.Append()
}

func (e *StatementAdd) Remove() {
	e.block.Remove()
}

func (e *StatementAdd) SetMainSvg(svg *html.TagSvg) {
	e.block.SetMainSvg(svg)
}

func (e *StatementAdd) SetResizerButton(resizeButton block.ResizeButton) {
	e.block.SetResizerButton(resizeButton)
}

func (e *StatementAdd) SetDraggerButton(draggerButton block.ResizeButton) {
	e.block.SetDraggerButton(draggerButton)
}

func (e *StatementAdd) SetGridAdjust(gridAdjust rulesStage.GridAdjust) {
	e.gridAdjust = gridAdjust
	e.block.SetGridAdjust(gridAdjust)
}

// SetWarning sets the visibility of the warning mark
func (e *StatementAdd) SetWarning(warning bool) {
	if !e.block.GetInitialized() {
		return
	}

	e.block.SetWarning(warning)
}

func (e *StatementAdd) SelectedInvert() {
	e.block.SelectedInvert()
}

func (e *StatementAdd) Get() (container *html.TagSvg) {
	return e.block.GetIdeStage()
}

func (e *StatementAdd) SetFatherId(fatherId string) {
	e.block.SetFatherId(fatherId)
}

func (e *StatementAdd) SetName(name string) {
	e.block.SetName(name)
}

func (e *StatementAdd) SetPosition(x, y rulesDensity.Density) {
	e.block.SetPosition(x, y)
}

func (e *StatementAdd) SetSize(wight, height rulesDensity.Density) {
	e.block.SetSize(wight, height)
}

func (e *StatementAdd) getMenuLabel(condition bool, labelTrue, labelFalse string) (label string) {
	if condition {
		return labelTrue
	}

	return labelFalse
}

func (e *StatementAdd) getMenu() (content []components.MenuOptions) {
	content = []components.MenuOptions{
		{
			Type: "grid",
			Items: []components.MenuOptions{
				{
					Label:  "Cat 1",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
				},
				{
					Label:  "Cat 2",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
				},
				{
					Label:  "Cat 3",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 3"); return nil }),
				},
				{
					Label:  "Cat 4",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 4"); return nil }),
				},
				{
					Label:  "Cat 5",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 5"); return nil }),
				},
				{
					Label:  "Cat 6",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 6"); return nil }),
				},
			},
		},
		{
			Label: "-",
		},
		{
			Label: "Label 1",
			//Icon:      "icon 1",
			//IconLeft:  "icon left 1",
			//IconRight: "icon right 1",
			Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 1 ok"); return nil }),
		},
		{
			Label: "Label 2",
			//Icon:      "icon 1",
			//IconLeft:  "icon left 1",
			//IconRight: "icon right 1",
			Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 2 ok"); return nil }),
			Submenu: []components.MenuOptions{
				{
					Label: "Label 1",
					//Icon:      "icon 1",
					//IconLeft:  "icon left 1",
					//IconRight: "icon right 1",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 1 ok"); return nil }),
				},
				{
					Label: "Label 2",
					//Icon:      "icon 1",
					//IconLeft:  "icon left 1",
					//IconRight: "icon right 1",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 2 ok"); return nil }),
					Submenu: []components.MenuOptions{
						{
							Type: "grid",
							Items: []components.MenuOptions{
								{
									Label:  "Cat 1",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
								{
									Label:  "Cat 2",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
								},
								{
									Label:  "Cat 3",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 3"); return nil }),
								},
								{
									Label:  "Cat 4",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 4"); return nil }),
								},
								{
									Label:  "Cat 5",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 5"); return nil }),
								},
								{
									Label:  "Cat 6",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 6"); return nil }),
									Submenu: []components.MenuOptions{
										{
											Type: "grid",
											Items: []components.MenuOptions{
												{
													Label:  "Cat 1",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
												},
												{
													Label:  "Cat 2",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
												},
												{
													Label:  "Cat 3",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 3"); return nil }),
												},
												{
													Label:  "Cat 4",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 4"); return nil }),
												},
												{
													Label:  "Cat 5",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 5"); return nil }),
												},
												{
													Label:  "Cat 6",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 6"); return nil }),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Type: "grid",
			Items: []components.MenuOptions{
				{
					Label:  "Cat 1",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
				},
				{
					Label:  "Cat 2",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
				},
				{
					Label:  "Cat 3",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 3"); return nil }),
				},
				{
					Label:  "Cat 4",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 4"); return nil }),
				},
				{
					Label:  "Cat 5",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 5"); return nil }),
				},
				{
					Label:  "Cat 6",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 6"); return nil }),
				},
			},
		},
	}

	content = append(e.block.GetMenuDebug(), content...)

	return
}

func (e *StatementAdd) Init() (err error) {
	e.SetFatherId(rulesStage.KStageId)
	e.SetName("stmAdd")

	warningMark := new(ornament.WarningMarkExclamation)
	warningMark.SetMargin(0)
	_ = warningMark.Init()
	e.block.SetWarningMark(warningMark)

	size := rulesDensity.Density(60)
	e.defaultWidth = size
	e.defaultHeight = size
	e.horizontalMinimumSize = size
	e.verticalMinimumSize = size

	if e.block.GetWidth() == 0 {
		e.block.SetWidth(e.defaultWidth)
	}

	if e.block.GetHeight() == 0 {
		e.block.SetHeight(e.defaultHeight)
	}

	//e.block.SetName("StatementAdd")

	e.id = rulesSequentialId.GetIdFromBase(e.block.GetName())

	//e.block.SetDrag(true)
	e.block.SetResizeLocked(true)
	//e.block.SetEnableResize(true)
	//e.block.SetSelected(true)
	e.block.SetMinimumWidth(e.horizontalMinimumSize)
	e.block.SetMinimumHeight(e.verticalMinimumSize)

	e.ornamentDraw = new(math.OrnamentAdd)
	e.ornamentDrawIcon = new(math.OrnamentAdd)

	inputXSetup := connection.Setup{
		FatherId:           e.id,
		Name:               "inputX",
		DataType:           "int",
		AcceptNotConnected: false,
		LookedUp:           false,
		IsADataInput:       true,
		ClickFunc: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := this.Call("getConnData")
			log.Printf("FatherId: %v", data.Get("FatherId").String())
			log.Printf("Name: %v", data.Get("Name").String())
			log.Printf("DataType: %v", data.Get("DataType").String())
			log.Printf("AcceptNotConnected: %v", data.Get("AcceptNotConnected").Bool())
			log.Printf("LookedUp: %v", data.Get("LookedUp").Bool())
			log.Printf("IsADataInput: %v", data.Get("IsADataInput").Bool())
			log.Printf("Top: %v", data.Get("Top").Int())
			log.Printf("Left: %v", data.Get("Left").Int())
			return nil
		}),
	}
	if err = inputXSetup.Verify(); err != nil {
		log.Printf("stopButton.Verify: %v", err)
		return
	}
	e.ornamentDraw.InputXSetup(inputXSetup)

	inputYSetup := connection.Setup{
		FatherId:           e.id,
		Name:               "inputY",
		DataType:           "int",
		AcceptNotConnected: false,
		LookedUp:           false,
		IsADataInput:       true,
		ClickFunc: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := this.Call("getConnData")
			log.Printf("FatherId: %v", data.Get("FatherId").String())
			log.Printf("Name: %v", data.Get("Name").String())
			log.Printf("DataType: %v", data.Get("DataType").String())
			log.Printf("AcceptNotConnected: %v", data.Get("AcceptNotConnected").Bool())
			log.Printf("LookedUp: %v", data.Get("LookedUp").Bool())
			log.Printf("IsADataInput: %v", data.Get("IsADataInput").Bool())
			log.Printf("Top: %v", data.Get("Top").Int())
			log.Printf("Left: %v", data.Get("Left").Int())
			return nil
		}),
	}
	if err = inputYSetup.Verify(); err != nil {
		log.Printf("stopButton.Verify: %v", err)
		return
	}
	e.ornamentDraw.InputYSetup(inputYSetup)

	outputSetup := connection.Setup{
		FatherId:           e.id,
		Name:               "output",
		DataType:           "int",
		AcceptNotConnected: false,
		LookedUp:           false,
		IsADataInput:       false,
		ClickFunc: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := this.Call("getConnData")
			log.Printf("FatherId: %v", data.Get("FatherId").String())
			log.Printf("Name: %v", data.Get("Name").String())
			log.Printf("DataType: %v", data.Get("DataType").String())
			log.Printf("AcceptNotConnected: %v", data.Get("AcceptNotConnected").Bool())
			log.Printf("LookedUp: %v", data.Get("LookedUp").Bool())
			log.Printf("IsADataInput: %v", data.Get("IsADataInput").Bool())
			log.Printf("Top: %v", data.Get("Top").Int())
			log.Printf("Left: %v", data.Get("Left").Int())
			return nil
		}),
	}
	if err = outputSetup.Verify(); err != nil {
		log.Printf("stopButton.Verify: %v", err)
		return
	}
	e.ornamentDraw.OutputSetup(outputSetup)

	_ = e.ornamentDraw.Init()
	_ = e.ornamentDrawIcon.Init()

	e.block.SetOrnament(e.ornamentDraw)

	_ = e.block.Init()

	if err = e.block.SetID(e.id); err != nil {
		log.Printf("e.block.SetID(e.id): %s", err)
		return
	}

	//e.block.SetOnResize(func(element js.Value, width, height int) {
	//	e.inputA.SetX(width - 50 - 4)
	//	e.inputA.SetY(height - 40 - 2)
	//})

	e.menu.SetNode(e.block.GetDeviceDiv())
	e.menu.SetTitle("Add Statement")
	e.menu.SetContentFunc(e.getMenu)
	e.menu.Init()

	return nil
}

func (e *StatementAdd) GetIconName() (name string) {
	return "Add"
}

func (e *StatementAdd) GetIconCategory() (name string) {
	return "Math"
}

//func (e *StatementAdd) GetIcon() (icon []js.Value) {
//	xc := rulesIcon.Width / rulesDensity.Density(4)
//	yc := rulesIcon.Height * rulesDensity.Density(0.15)
//	wOrn := rulesIcon.Width / rulesDensity.Density(2)
//	ornamentSvg := e.ornamentDrawIcon.
//		GetSvg().
//		X(xc.GetInt() + rulesDensity.Density(5).GetInt()).
//		Y(yc.GetInt())
//
//	_ = e.ornamentDrawIcon.Update(0, 0, wOrn, wOrn)
//
//	path := utilsDraw.PolygonPath(6, wOrn, wOrn, wOrn, 0)
//	iconPath := factoryBrowser.NewTagSvgPath().
//		StrokeWidth(rulesIcon.BorderWidth.GetInt()).
//		Stroke(rulesIcon.BorderColor).
//		Fill(rulesIcon.FillColor).
//		D(path)
//
//	iconText := factoryBrowser.NewTagSvgText().
//		X(rulesDensity.Density(85).GetInt()).
//		Y(rulesIcon.TextY.GetInt()).
//		Text("Add").
//		Fill(rulesIcon.TextColor).
//		FontFamily(rulesIcon.FontFamily).
//		//Filter("url(#textBlur)").
//		FontSize(rulesIcon.FontSize.GetInt())
//
//	iconSvg := factoryBrowser.NewTagSvg().
//		Width(rulesIcon.Width.GetFloat()).
//		Height(rulesIcon.Height.GetFloat()).
//		Append(iconPath, ornamentSvg, iconText, rulesIcon.FilterIcon, rulesIcon.FilterText)
//
//	w := rulesIcon.Width * rulesIcon.SizeRatio
//	h := rulesIcon.Height * rulesIcon.SizeRatio
//	return iconSvg.ToCanvas(w.GetFloat(), h.GetFloat())
//}

func (e *StatementAdd) GetInitialized() (initialized bool) {
	return e.block.GetInitialized()
}

func (e *StatementAdd) GetWarning() (warning bool) {
	return e.block.GetWarning()
}

func (e *StatementAdd) GetDragBlocked() (blocked bool) {
	return e.block.GetDragLocked()
}

func (e *StatementAdd) GetDragEnable() (enabled bool) {
	return e.block.GetDragEnable()
}

func (e *StatementAdd) GetResize() (enabled bool) {
	return e.block.GetResizeEnable()
}

func (e *StatementAdd) GetResizeBlocked() (blocked bool) {
	return e.block.GetResizeLocked()
}

func (e *StatementAdd) GetSelectBlocked() (blocked bool) {
	return e.block.GetSelectLocked()
}

func (e *StatementAdd) GetSelected() (selected bool) {
	return e.block.GetSelected()
}

func (e *StatementAdd) GetID() (id string) {
	return e.block.GetID()
}

func (e *StatementAdd) GetName() (name string) {
	return e.block.GetName()
}

func (e *StatementAdd) GetWidth() (width rulesDensity.Density) {
	return e.block.GetWidth()
}

func (e *StatementAdd) GetHeight() (height rulesDensity.Density) {
	return e.block.GetHeight()
}

func (e *StatementAdd) GetX() (x rulesDensity.Density) {
	return e.block.GetX()
}

func (e *StatementAdd) GetY() (y rulesDensity.Density) {
	return e.block.GetY()
}

func (e *StatementAdd) SetX(x rulesDensity.Density) {
	e.block.SetX(x)
}

func (e *StatementAdd) SetY(y rulesDensity.Density) {
	e.block.SetY(y)
}

func (e *StatementAdd) SetWidth(width rulesDensity.Density) {
	e.block.SetWidth(width)
}

func (e *StatementAdd) SetHeight(height rulesDensity.Density) {
	e.block.SetHeight(height)
}

func (e *StatementAdd) SetSelected(selected bool) {
	e.block.SetSelected(selected)
}

func (e *StatementAdd) SetDragEnable(enabled bool) {
	e.block.SetDragEnable(enabled)
}

func (e *StatementAdd) GetResizeEnable() (enabled bool) {
	return e.block.GetResizeEnable()
}

func (e *StatementAdd) SetResizeEnable(enabled bool) {
	e.block.SetResizeEnable(enabled)
}

func (e *StatementAdd) GetZIndex() (zIndex int) {
	return e.block.GetZIndex()
}
