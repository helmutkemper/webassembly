package devices

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/connection"
	"github.com/helmutkemper/webassembly/examples/ide/devices/block"
	"github.com/helmutkemper/webassembly/examples/ide/ornament/doubleLoopArrow"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"log"
	"reflect"
	"syscall/js"
)

type StatementLoop struct {
	block block.Block
	menu  Menu

	debugSelected bool

	defaultWidth          int
	defaultHeight         int
	horizontalMinimumSize int
	verticalMinimumSize   int
	ornamentDraw          *doubleLoopArrow.DoubleLoopArrow
	id                    string
	connStop              *connection.Connection
	debugMode             bool

	sequentialId utils.SequentialInterface
}

func (e *StatementLoop) Get() (container *html.TagDiv) {
	return e.block.GetIdeStage()
}

func (e *StatementLoop) SetFatherId(fatherId string) {
	e.block.SetFatherId(fatherId)
}

func (e *StatementLoop) SetName(name string) (err error) {
	return e.block.SetName(name)
}

func (e *StatementLoop) SetPosition(x, y int) {
	e.block.SetPosition(x, y)
}

func (e *StatementLoop) SetSize(wight, height int) {
	e.block.SetSize(wight, height)
}

func (e *StatementLoop) SetSequentialId(sequentialId utils.SequentialInterface) {
	e.sequentialId = sequentialId
}

func (e *StatementLoop) getMenuLabel(condition bool, labelTrue, labelFalse string) (label string) {
	if condition {
		return labelTrue
	}

	return labelFalse
}

func (e *StatementLoop) getMenu() (content []components.MenuOptions) {
	content = []components.MenuOptions{
		{
			Label: "Debug",
			Submenu: []components.MenuOptions{
				{
					Label: e.getMenuLabel(e.debugSelected, "Unselect", "Select"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.debugSelected = !e.debugSelected
						e.block.SetSelected(e.debugSelected)
						e.menu.ReInit()

						return nil
					}),
				},
				{
					Label:  "Resize",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { e.block.SetResize(true); return nil }),
				},
			},
		},
		{
			Label: "Debug",
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

	return
}

func (e *StatementLoop) Init() (err error) {
	e.defaultWidth = 500
	e.defaultHeight = 400
	e.horizontalMinimumSize = 400
	e.verticalMinimumSize = 300

	if e.block.GetWidth() == 0 {
		e.block.SetWidth(e.defaultWidth)
	}

	if e.block.GetHeight() == 0 {
		e.block.SetHeight(e.defaultHeight)
	}

	err = e.block.SetName("StatementLoop")
	if err != nil {
		return
	}

	e.id, err = e.sequentialId.GetId(e.block.GetName())
	if err != nil {
		return
	}

	e.block.SetDragEnabled(true)
	//e.block.SetEnableResize(true)
	//e.block.SetSelected(true)
	e.block.SetMinimumWidth(e.horizontalMinimumSize)
	e.block.SetMinimumHeight(e.verticalMinimumSize)

	e.ornamentDraw = new(doubleLoopArrow.DoubleLoopArrow)
	e.ornamentDraw.SetWarningMarkMargin(20)
	_ = e.ornamentDraw.Init()

	e.block.SetOrnament(e.ornamentDraw)

	_ = e.block.Init()

	if err = e.block.SetID(e.id); err != nil {
		return
	}

	e.connStop = new(connection.Connection)
	e.connStop.Create(e.block.GetWidth()-50-4, e.block.GetHeight()-40-2, 4, 3, factoryColor.NewRed())
	e.connStop.SetFather(e.block.GetDeviceDiv())
	e.connStop.SetAsInput()
	_ = e.connStop.SetName("stop")
	e.connStop.SetDataType(reflect.Bool)
	e.connStop.SetAcceptedNotConnected(true)
	e.connStop.SetBlocked(false)
	e.connStop.Init()

	e.block.SetOnResize(func(element js.Value, width, height int) {
		e.connStop.SetX(width - 50 - 4)
		e.connStop.SetY(height - 40 - 2)
	})

	e.menu.SetNode(e.block.GetDeviceDiv())
	e.menu.SetTitle("Loop")
	e.menu.SetContentFunc(e.getMenu)
	e.menu.Init()

	return nil
}
