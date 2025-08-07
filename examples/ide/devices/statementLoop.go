package devices

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/connection"
	"github.com/helmutkemper/webassembly/examples/ide/devices/block"
	"github.com/helmutkemper/webassembly/examples/ide/manager"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/ornament/doubleLoopArrow"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/examples/ide/rulesIcon"
	"github.com/helmutkemper/webassembly/examples/ide/rulesSequentialId"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
	"github.com/helmutkemper/webassembly/examples/ide/translate"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/utilsDraw"
	"github.com/helmutkemper/webassembly/utilsText"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
	"syscall/js"
)

type StatementLoop struct {
	block block.Block
	menu  Menu

	debugSelected bool

	defaultWidth          rulesDensity.Density
	defaultHeight         rulesDensity.Density
	horizontalMinimumSize rulesDensity.Density
	verticalMinimumSize   rulesDensity.Density
	ornamentDraw          *doubleLoopArrow.DoubleLoopArrow
	ornamentDrawIcon      *doubleLoopArrow.DoubleLoopArrow
	id                    string
	//connStop              *connection.Connection
	debugMode bool

	gridAdjust rulesStage.GridAdjust

	iconStatus int
}

func (e *StatementLoop) Append() {
	e.block.Append()
}

func (e *StatementLoop) Remove() {
	e.block.Remove()
}

func (e *StatementLoop) SetMainSvg(svg *html.TagSvg) {
	e.block.SetMainSvg(svg)
}

func (e *StatementLoop) SetResizerButton(resizeButton block.ResizeButton) {
	e.block.SetResizerButton(resizeButton)
}

func (e *StatementLoop) SetDraggerButton(draggerButton block.ResizeButton) {
	e.block.SetDraggerButton(draggerButton)
}

func (e *StatementLoop) SetGridAdjust(gridAdjust rulesStage.GridAdjust) {
	e.gridAdjust = gridAdjust
	e.block.SetGridAdjust(gridAdjust)
}

func (e *StatementLoop) GetWidth() (width rulesDensity.Density) {
	return e.block.GetWidth()
}

func (e *StatementLoop) GetHeight() (height rulesDensity.Density) {
	return e.block.GetHeight()
}

// SetWarning sets the visibility of the warning mark
func (e *StatementLoop) SetWarning(warning bool) {
	if !e.block.GetInitialized() {
		log.Println("Warning: block is't initialized")
		return
	}

	e.block.SetWarning(warning)
}

func (e *StatementLoop) Get() (container *html.TagSvg) {
	return e.block.GetIdeStage()
}

func (e *StatementLoop) SetFatherId(fatherId string) {
	e.block.SetFatherId(fatherId)
}

func (e *StatementLoop) SetName(name string) {
	e.block.SetName(name)
}

func (e *StatementLoop) SetPosition(x, y rulesDensity.Density) {
	e.block.SetPosition(x, y)
}

func (e *StatementLoop) SetSize(wight, height rulesDensity.Density) {
	e.block.SetSize(wight, height)
}

func (e *StatementLoop) getMenuLabel(condition *bool, labelTrue, labelFalse string) (label string) {
	defer func() {
		*condition = !*condition
	}()

	if *condition {
		return labelTrue
	}

	return labelFalse
}

func (e *StatementLoop) getMenu() (content []components.MenuOptions) {
	content = []components.MenuOptions{
		{
			Type: "grid",
			Items: []components.MenuOptions{
				{
					Label: "Cat 1",
					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					//Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
					Submenu: e.block.GetMenuDebug(),
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

	//content = append(e.block.GetMenuDebug(), content...)
	return
}

//func (e *StatementLoop) SetSelected(selected bool) {
//	e.block.SetSelected(selected)
//	e.ornamentDraw.SetSelected(selected)
//}

func (e *StatementLoop) Init() (err error) {

	warningMark := new(ornament.WarningMarkExclamation)
	warningMark.SetMargin(15)
	_ = warningMark.Init()
	e.block.SetWarningMark(warningMark)

	e.SetFatherId(rulesStage.KStageId)
	e.SetName("stmLoop")

	e.defaultWidth = 400
	e.defaultHeight = 300

	defaultWidth, defaultHeight := e.gridAdjust.AdjustCenter(e.defaultWidth.GetInt(), e.defaultHeight.GetInt())
	e.defaultWidth, e.defaultHeight = rulesDensity.Convert(defaultWidth), rulesDensity.Convert(defaultHeight)

	e.horizontalMinimumSize = 150
	e.verticalMinimumSize = 150

	horizontalMinimumSize, verticalMinimumSize := e.gridAdjust.AdjustCenter(e.horizontalMinimumSize.GetInt(), e.verticalMinimumSize.GetInt())
	e.horizontalMinimumSize, e.verticalMinimumSize = rulesDensity.Convert(horizontalMinimumSize), rulesDensity.Convert(verticalMinimumSize)

	if e.block.GetWidth() == 0 || e.block.GetHeight() == 0 {
		e.block.SetSize(e.defaultWidth, e.defaultHeight)
	}

	//e.block.SetName("StatementLoop")
	e.id = rulesSequentialId.GetIdFromBase(e.block.GetName())

	//e.block.SetDrag(true)
	//e.block.SetEnableResize(true)
	//e.block.SetSelected(true)
	e.block.SetMinimumWidth(e.horizontalMinimumSize)
	e.block.SetMinimumHeight(e.verticalMinimumSize)

	e.ornamentDraw = new(doubleLoopArrow.DoubleLoopArrow)
	e.ornamentDrawIcon = new(doubleLoopArrow.DoubleLoopArrow)

	stopButton := connection.Setup{
		FatherId:           e.id,
		Name:               "stopButton",
		DataType:           "bool",
		AcceptNotConnected: true,
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
	if err = stopButton.Verify(); err != nil {
		log.Printf("stopButton.Verify: %v", err)
		return
	}

	e.ornamentDraw.StopButtonSetup(stopButton)

	if err = e.ornamentDraw.GetConnectionError(); err != nil {
		return
	}

	_ = e.ornamentDraw.Init()
	_ = e.ornamentDrawIcon.Init()

	e.block.SetOrnament(e.ornamentDraw)
	_ = e.block.Init()

	if err = e.block.SetID(e.id); err != nil {
		return
	}

	//e.connStop = new(connection.Connection)
	//e.connStop.Create(e.block.GetWidth()-57, e.block.GetHeight()-42)
	//e.connStop.SetFather(e.block.GetDeviceDiv())
	//e.connStop.SetAsInput()
	//_ = e.connStop.SetName("stop")
	//e.connStop.SetDataType(reflect.Bool)
	//e.connStop.SetAcceptedNotConnected(true)
	//e.connStop.SetBlocked(false)
	//e.connStop.Init()
	//
	//e.block.SetOnResize(func(element js.Value, width, height int) {
	//	e.connStop.SetX(width - 57)
	//	e.connStop.SetY(height - 42)
	//})

	//e.ornamentDraw.SetStopButtonMouseClick(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	//	log.Printf("name: %v", this.Get("dataset").Get("name"))
	//	log.Printf("type: %v", this.Get("dataset").Get("type"))
	//	log.Printf("notConnected: %v", this.Get("dataset").Get("notConnected"))
	//	log.Printf("locked: %v", this.Get("dataset").Get("locked"))
	//	log.Printf("fatherId: %v", this.Get("dataset").Get("fatherId"))
	//	log.Printf("direction: %v", this.Get("dataset").Get("direction"))
	//	return nil
	//}))

	e.menu.SetNode(e.block.GetDeviceDiv())
	e.menu.SetTitle("Loop")
	e.menu.SetContentFunc(e.getMenu)
	e.menu.Init()

	//e.block.SetResizeEnable(true)
	//e.block.SetSelected(true)

	return nil
}

func (e *StatementLoop) GetIconName() (name string) {
	return "Loop"
}

func (e *StatementLoop) GetIconCategory() (name string) {
	return "Loop"
}

func (e *StatementLoop) SetStatus(status int) {
	e.iconStatus = status
}

func (e *StatementLoop) GetStatus() (staus int) {
	return e.iconStatus
}

func (e *StatementLoop) GetIcon() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconDeviceLoop",
				Other: "Loop",
			},
		},
	)
	if err != nil {
		translated = "Loop"
		log.Printf("icon translation error: %v", err)
	}

	name := e.GetIconName()
	category := e.GetIconCategory()
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Name:     name,
			Category: category,
			Label:    translated,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Name:     name,
			Category: category,
			Label:    translated,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Name:     name,
			Category: category,
			Label:    translated,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Name:     name,
			Category: category,
			Label:    translated,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Name:     name,
			Category: category,
			Label:    translated,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e *StatementLoop) getIcon(data rulesIcon.Data) (png js.Value) {

	data = rulesIcon.DataVerifyElementIcon(data)

	// icon body
	svgIcon := factoryBrowser.NewTagSvg().
		X(rulesIcon.Width.GetInt() / 2).
		Y(rulesIcon.Height.GetInt() / 2).
		Width(rulesIcon.Width.GetInt()).
		Height(rulesIcon.Height.GetInt())

	// hexagon maker
	hexPath := utilsDraw.PolygonPath(
		6,
		rulesIcon.Width/2,
		rulesIcon.Width/2,
		rulesIcon.Width/2,
		0,
	)

	// svg hexagon
	hexDraw := factoryBrowser.NewTagSvgPath().
		StrokeWidth(rulesIcon.BorderWidth.GetInt()).
		Stroke(data.ColorBorder).
		Fill(data.ColorBackground).
		D(hexPath)

	xc := rulesIcon.Width / 4
	yc := rulesIcon.Height * 0.15
	wOrn := rulesIcon.Width / 2
	icon := e.ornamentDrawIcon.
		GetSvg().
		X(xc.GetInt()).
		Y(yc.GetInt())

	_ = e.ornamentDrawIcon.Update(0, 0, wOrn, wOrn)

	// calc width label
	widthLabel, _ := utilsText.GetTextSize(
		data.Label,
		rulesIcon.FontFamily,
		rulesIcon.FontWeight,
		rulesIcon.FontStyle,
		data.LabelFontSize.GetInt(),
	)

	// label, svg text
	label := factoryBrowser.NewTagSvgText().
		FontFamily(rulesIcon.FontFamily).
		FontWeight(rulesIcon.FontWeight).
		FontStyle(rulesIcon.FontStyle).
		FontSize(data.LabelFontSize.GetInt()).
		Text(data.Label).
		Fill(data.ColorLabel).
		X((rulesIcon.Width / 2).GetInt() - widthLabel/2).
		Y(data.LabelY.GetInt())
	svgIcon.Append(hexDraw, icon, label)

	w := rulesIcon.Width * rulesIcon.SizeRatio
	h := rulesIcon.Height * rulesIcon.SizeRatio
	return svgIcon.ToCanvas(
		html.CanvasData{
			Width:  w.GetInt(),
			Height: h.GetInt(),
		},
	)
}

func (e *StatementLoop) GetInitialized() (initialized bool) {
	return e.block.GetInitialized()
}

func (e *StatementLoop) GetWarning() (warning bool) {
	return e.block.GetWarning()
}

func (e *StatementLoop) GetDragBlocked() (blocked bool) {
	return e.block.GetDragLocked()
}

func (e *StatementLoop) GetDragEnable() (enabled bool) {
	return e.block.GetDragEnable()
}

func (e *StatementLoop) GetResize() (enabled bool) {
	return e.block.GetResizeEnable()
}

func (e *StatementLoop) GetResizeBlocked() (blocked bool) {
	return e.block.GetResizeLocked()
}

func (e *StatementLoop) GetSelectBlocked() (blocked bool) {
	return e.block.GetSelectLocked()
}

func (e *StatementLoop) GetSelected() (selected bool) {
	return e.block.GetSelected()
}

func (e *StatementLoop) onConnectionClick() {}

func (e *StatementLoop) GetID() (id string) {
	return e.block.GetID()
}

func (e *StatementLoop) GetName() (name string) {
	return e.block.GetName()
}

func (e *StatementLoop) GetX() (x rulesDensity.Density) {
	return e.block.GetX()
}

func (e *StatementLoop) GetY() (y rulesDensity.Density) {
	return e.block.GetY()
}

func (e *StatementLoop) SetX(x rulesDensity.Density) {
	e.block.SetX(x)
}

func (e *StatementLoop) SetY(y rulesDensity.Density) {
	e.block.SetY(y)
}

func (e *StatementLoop) SetWidth(width rulesDensity.Density) {
	e.block.SetWidth(width)
}

func (e *StatementLoop) SetHeight(height rulesDensity.Density) {
	e.block.SetHeight(height)
}

func (e *StatementLoop) SetSelected(selected bool) {
	e.block.SetSelected(selected)
}

func (e *StatementLoop) SetDragEnable(enabled bool) {
	e.block.SetDragEnable(enabled)
}

func (e *StatementLoop) GetResizeEnable() (enabled bool) {
	return e.block.GetResizeEnable()
}

func (e *StatementLoop) SetResizeEnable(enabled bool) {
	e.block.SetResizeEnable(enabled)
}

func (e *StatementLoop) GetZIndex() (zIndex int) {
	return e.block.GetZIndex()
}
