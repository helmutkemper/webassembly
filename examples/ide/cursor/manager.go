package cursor

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"image/color"
	"log"
	"syscall/js"
)

type manager struct {
	cursor CursorControl

	lastPointType string

	id       string
	block    *html.TagDiv
	ideStage *html.TagSvg

	listenerClick js.Func
	listenerMove  js.Func

	coordinates []coordinate

	bindingType string
	fatherId    string
}

func (e *manager) getLastPointType() (pointType string) {
	if e.lastPointType == "vertical" {
		e.lastPointType = "horizontal"

		return e.lastPointType
	}

	e.lastPointType = "horizontal"
	return e.lastPointType
}

func (e *manager) AddPoint(x, y int, t string) {

	if t == "" {
		t = e.getLastPointType()
	}

	e.coordinates = append(
		e.coordinates,
		coordinate{
			x: x,
			y: y,
			t: t,
		},
	)
}

func (e *manager) ActivateVertical(x, y int, color color.RGBA) {
	e.lastPointType = "vertical"
	e.AddPoint(x, y, "vertical")

	e.initEvents()

	e.cursor.SetColor(color)
	e.cursor.ShowHorizontal()
}

func (e *manager) ActivateHorizontal(x, y int, color color.RGBA) {
	e.lastPointType = "horizontal"
	e.AddPoint(x, y, "horizontal")

	e.initEvents()

	e.cursor.SetColor(color)
	e.cursor.ShowVertical()
}

func (e *manager) Deactivate() {
	e.removeEvents()
}

// SetOrnament Sets the ornament draw object
//
//	ornament draw object is the instance in charge of making the SVG design of the device
func (e *manager) SetOrnament(ornament CursorControl) {
	//e.ornament = ornament
}

// SetFatherId Receives the div ID used as a stage for the IDE and puts it to occupy the entire browser area
func (e *manager) SetFatherId(fatherId string) {
	e.fatherId = fatherId
	e.ideStage = factoryBrowser.NewTagSvg().
		Import(fatherId) //.
	//AddStyle("position", "relative").
	//AddStyle("width", "100vw"). // todo: transformar isto em algo chamado uma única vez e em outro lugar
	//AddStyle("height", "100vh") // todo: transformar isto em algo chamado uma única vez e em outro lugar
}

func (e *manager) initEvents() {
	e.initEventMouseMove()
	e.initEventMouseClick()
}

func (e *manager) initEventMouseClick() {
	document := js.Global().Get("document")
	document.Call("addEventListener", "click", e.listenerClick)
}

func (e *manager) initEventMouseMove() {
	document := js.Global().Get("document")
	document.Call("addEventListener", "mousemove", e.listenerMove)
}

func (e *manager) removeEvents() {
	document := js.Global().Get("document")
	document.Call("removeEventListener", "click", e.listenerClick)
	document.Call("removeEventListener", "mousemove", e.listenerMove)
}

func (e *manager) getBindingType() (binding string) {
	if e.bindingType == "horizontal" {
		e.bindingType = "vertical"
	} else {
		e.bindingType = "horizontal"
	}

	return e.bindingType
}

func (e *manager) updateOrnament() (err error) {
	width := e.block.GetClientWidth()
	height := e.block.GetClientHeight()
	_ = e.cursor.Update(width, height)
	//_ = e.ornament.Update(width, height)
	return
}

func (e *manager) Init() {
	e.cursor = new(cursor) // todo: mudar daqui

	e.coordinates = make([]coordinate, 0)

	e.listenerClick = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		element := args[0]
		x := element.Get("clientX").Int()
		y := element.Get("clientY").Int()
		log.Println("x:", x, "y:", y)
		e.cursor.ChangesDirection()
		//e.ornament.AddPoint(x, y, e.getBindingType())
		e.updateOrnament()
		return nil
	})

	e.listenerMove = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		element := args[0]
		x := element.Get("clientX").Int()
		y := element.Get("clientY").Int()
		width := e.block.GetClientWidth()
		height := e.block.GetClientHeight()
		e.cursor.Move(x, y, width, height)

		//e.ornament.Cursor(x, y, width, height)
		return nil
	})

	_ = e.createBlock()
	e.initEvents()

	e.bindingType = "horizontal"

	if e.cursor != nil {
		e.cursor.SetFatherId(e.fatherId)
		e.cursor.Init()
		svg := e.cursor.GetSvg()
		e.block.Append(svg)
		e.updateOrnament()
		e.cursor.ShowVertical()
		e.cursor.Hide()
	}
}

func (e *manager) createBlock() (err error) {
	e.id = "connections" //todo: fazer

	e.block = factoryBrowser.NewTagDiv().
		Id(e.id).
		//Class(e.classListName).
		AddStyle("position", "absolute").
		AddStyle("top", fmt.Sprintf("%dpx", 0)).
		AddStyle("left", fmt.Sprintf("%dpx", 0)).
		AddStyle("width", "100%").
		AddStyle("height", "100%")
	e.ideStage.Append(e.block)

	return
}
