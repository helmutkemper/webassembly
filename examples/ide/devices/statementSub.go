package devices

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/connection"
	"github.com/helmutkemper/webassembly/examples/ide/devices/block"
	"github.com/helmutkemper/webassembly/examples/ide/manager"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/ornament/math"
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

type StatementSub struct {
	block block.Block
	menu  Menu

	defaultWidth          rulesDensity.Density
	defaultHeight         rulesDensity.Density
	horizontalMinimumSize rulesDensity.Density
	verticalMinimumSize   rulesDensity.Density

	ornamentDraw     *math.OrnamentSubtract
	ornamentDrawIcon *math.OrnamentSubtract

	id string
	//inputA                *connection.Connection
	//inputB                *connection.Connection
	//output                *connection.Connection
	debugMode bool

	gridAdjust rulesStage.GridAdjust

	iconStatus int
}

func (e *StatementSub) Append() {
	e.block.Append()
}

func (e *StatementSub) Remove() {
	e.block.Remove()
}

func (e *StatementSub) SetMainSvg(svg *html.TagSvg) {
	e.block.SetMainSvg(svg)
}

func (e *StatementSub) SetResizerButton(resizeButton block.ResizeButton) {
	e.block.SetResizerButton(resizeButton)
}

func (e *StatementSub) SetDraggerButton(draggerButton block.ResizeButton) {
	e.block.SetDraggerButton(draggerButton)
}

func (e *StatementSub) SetGridAdjust(gridAdjust rulesStage.GridAdjust) {
	e.gridAdjust = gridAdjust
	e.block.SetGridAdjust(gridAdjust)
}

// SetWarning sets the visibility of the warning mark
func (e *StatementSub) SetWarning(warning bool) {
	if !e.block.GetInitialized() {
		return
	}

	e.block.SetWarning(warning)
}

func (e *StatementSub) SelectedInvert() {
	e.block.SelectedInvert()
}

func (e *StatementSub) Get() (container *html.TagSvg) {
	return e.block.GetIdeStage()
}

func (e *StatementSub) SetFatherId(fatherId string) {
	e.block.SetFatherId(fatherId)
}

func (e *StatementSub) SetName(name string) {
	e.block.SetName(name)
}

func (e *StatementSub) SetPosition(x, y rulesDensity.Density) {
	e.block.SetPosition(x, y)
}

func (e *StatementSub) SetSize(wight, height rulesDensity.Density) {
	e.block.SetSize(wight, height)
}

func (e *StatementSub) getMenuLabel(condition bool, labelTrue, labelFalse string) (label string) {
	if condition {
		return labelTrue
	}

	return labelFalse
}

func (e *StatementSub) getMenu() (content []components.MenuOptions) {
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

func (e *StatementSub) Init() (err error) {
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

	e.ornamentDraw = new(math.OrnamentSubtract)
	e.ornamentDrawIcon = new(math.OrnamentSubtract)

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

func (e *StatementSub) GetIconName() (name string) {
	return "Sub"
}

func (e *StatementSub) GetIconCategory() (name string) {
	return "Math"
}

func (e *StatementSub) SetStatus(status int) {
	e.iconStatus = status
}

func (e *StatementSub) GetStatus() (staus int) {
	return e.iconStatus
}

func (e *StatementSub) GetIcon() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconDeviceSub",
				Other: "Sub",
			},
		},
	)
	if err != nil {
		translated = "Sub"
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

func (e *StatementSub) getIcon(data rulesIcon.Data) (png js.Value) {

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

	xc := rulesIcon.Width / rulesDensity.Density(4)
	yc := rulesIcon.Height * rulesDensity.Density(0.15)
	wOrn := rulesIcon.Width / rulesDensity.Density(2)
	icon := e.ornamentDrawIcon.
		GetSvg().
		X(xc.GetInt() + rulesDensity.Density(5).GetInt()).
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
	return svgIcon.ToCanvas(w.GetFloat(), h.GetFloat())
}

func (e *StatementSub) GetInitialized() (initialized bool) {
	return e.block.GetInitialized()
}

func (e *StatementSub) GetWarning() (warning bool) {
	return e.block.GetWarning()
}

func (e *StatementSub) GetDragBlocked() (blocked bool) {
	return e.block.GetDragLocked()
}

func (e *StatementSub) GetDragEnable() (enabled bool) {
	return e.block.GetDragEnable()
}

func (e *StatementSub) GetResize() (enabled bool) {
	return e.block.GetResizeEnable()
}

func (e *StatementSub) GetResizeBlocked() (blocked bool) {
	return e.block.GetResizeLocked()
}

func (e *StatementSub) GetSelectBlocked() (blocked bool) {
	return e.block.GetSelectLocked()
}

func (e *StatementSub) GetSelected() (selected bool) {
	return e.block.GetSelected()
}

func (e *StatementSub) GetID() (id string) {
	return e.block.GetID()
}

func (e *StatementSub) GetName() (name string) {
	return e.block.GetName()
}

func (e *StatementSub) GetWidth() (width rulesDensity.Density) {
	return e.block.GetWidth()
}

func (e *StatementSub) GetHeight() (height rulesDensity.Density) {
	return e.block.GetHeight()
}

func (e *StatementSub) GetX() (x rulesDensity.Density) {
	return e.block.GetX()
}

func (e *StatementSub) GetY() (y rulesDensity.Density) {
	return e.block.GetY()
}

func (e *StatementSub) SetX(x rulesDensity.Density) {
	e.block.SetX(x)
}

func (e *StatementSub) SetY(y rulesDensity.Density) {
	e.block.SetY(y)
}

func (e *StatementSub) SetWidth(width rulesDensity.Density) {
	e.block.SetWidth(width)
}

func (e *StatementSub) SetHeight(height rulesDensity.Density) {
	e.block.SetHeight(height)
}

func (e *StatementSub) SetSelected(selected bool) {
	e.block.SetSelected(selected)
}

func (e *StatementSub) SetDragEnable(enabled bool) {
	e.block.SetDragEnable(enabled)
}

func (e *StatementSub) GetResizeEnable() (enabled bool) {
	return e.block.GetResizeEnable()
}

func (e *StatementSub) SetResizeEnable(enabled bool) {
	e.block.SetResizeEnable(enabled)
}

func (e *StatementSub) GetZIndex() (zIndex int) {
	return e.block.GetZIndex()
}
