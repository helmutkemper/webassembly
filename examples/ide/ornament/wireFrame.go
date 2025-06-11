package ornament

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"log"
	"syscall/js"
)

type Frame interface {

	// Init Initializes the instance
	Init() (err error)

	// Update Draw the element design
	Update(width, height int) (err error)

	// GetSvg Returns the SVG tag with the element design
	GetSvg() (svg *html.TagSvg)

	// SetWarning sets the visibility of the warning mark
	SetWarning(warning bool)

	AddPoint(x, y int, bindingType string)

	MoveCursor(x, y, width, height int)
}

type Connections struct {
	id       string
	block    *html.TagDiv
	ideStage *html.TagDiv

	ornament Frame

	bindingType string
}

// SetOrnament Sets the ornament draw object
//
//	ornament draw object is the instance in charge of making the SVG design of the device
func (e *Connections) SetOrnament(ornament Frame) {
	e.ornament = ornament
}

// SetFatherId Receives the div ID used as a stage for the IDE and puts it to occupy the entire browser area
func (e *Connections) SetFatherId(fatherId string) {
	e.ideStage = factoryBrowser.NewTagDiv().
		Import(fatherId).
		AddStyle("position", "relative").
		AddStyle("width", "100vw"). // todo: transformar isto em algo chamado uma única vez e em outro lugar
		AddStyle("height", "100vh") // todo: transformar isto em algo chamado uma única vez e em outro lugar
}

func (e *Connections) initEvents() {
	document := js.Global().Get("document")
	document.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		element := args[0]
		x := element.Get("clientX").Int()
		y := element.Get("clientY").Int()
		log.Println("x:", x, "y:", y)
		e.ornament.AddPoint(x, y, e.getBindingType())
		e.updateOrnament()
		return nil
	}))

	document.Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		element := args[0]
		x := element.Get("clientX").Int()
		y := element.Get("clientY").Int()
		width := e.block.GetClientWidth()
		height := e.block.GetClientHeight()
		e.ornament.MoveCursor(x, y, width, height)
		return nil
	}))
}

func (e *Connections) getBindingType() (binding string) {
	if e.bindingType == "horizontal" {
		e.bindingType = "vertical"
	} else {
		e.bindingType = "horizontal"
	}

	return e.bindingType
}

func (e *Connections) updateOrnament() (err error) {
	width := e.block.GetClientWidth()
	height := e.block.GetClientHeight()
	_ = e.ornament.Update(width, height)
	return
}

func (e *Connections) Init() {
	_ = e.createBlock()
	e.initEvents()

	e.bindingType = "horizontal"

	if e.ornament != nil {
		svg := e.ornament.GetSvg()
		e.block.Append(svg)
		_ = e.updateOrnament()
	}
}

func (e *Connections) createBlock() (err error) {
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

type coordinate struct {
	x int
	y int
}

type WireFrame struct {
	spiderWay   []coordinate
	bindingType string

	svg       *html.TagSvg
	spiderWeb *html.TagSvgPath
	cursorWeb *html.TagSvgPath
}

func (e *WireFrame) BindingType(binding string) {

}

func (e *WireFrame) AddPoint(x, y int, bindingType string) {
	if e.spiderWay == nil {
		e.spiderWay = make([]coordinate, 0)
	}

	e.spiderWay = append(e.spiderWay, coordinate{x: x, y: y})
}

func (e *WireFrame) Init() (err error) {
	e.svg = factoryBrowser.NewTagSvg()

	e.spiderWeb = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke("red").
		StrokeWidth(3) //.
	//MarkerEnd("url(#spiderWeb)")
	e.svg.Append(e.spiderWeb)

	e.cursorWeb = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke("blue").
		StrokeWidth(3)
	e.svg.Append(e.cursorWeb)

	return
}

func (e *WireFrame) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

func (e *WireFrame) SetWarning(warning bool) {

}

func (e *WireFrame) MoveCursor(x, y, width, height int) {
	way := []string{
		fmt.Sprintf("M0,%v", y),
		fmt.Sprintf("L%v,%v", width, y),

		fmt.Sprintf("M%v,0", x),
		fmt.Sprintf("L%v,%v", x, height),
	}
	e.cursorWeb.D(way)
}

func (e *WireFrame) Update(width, height int) (err error) {
	e.svg.ViewBox([]int{0.0, 0.0, width, height})

	length := len(e.spiderWay)
	if length < 2 {
		return
	}

	way := []string{fmt.Sprintf("M %v %v", e.spiderWay[0].x, e.spiderWay[0].y)}
	for i := 1; i != length; i += 1 {
		way = append(way, fmt.Sprintf("L %v %v", e.spiderWay[i].x, e.spiderWay[i].y))
	}
	//way = append(way, "z")

	e.spiderWeb.D(way)

	return
}
