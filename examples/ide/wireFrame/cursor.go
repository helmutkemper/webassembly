package wireFrame

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
	"log"
	"syscall/js"
)

type cursor struct {
	ideStage   *html.TagDiv
	block      *html.TagDiv
	vertical   *html.TagSvgPath
	horizontal *html.TagSvgPath
	svg        *html.TagSvg
	visible    string
	document   js.Value
	body       js.Value
	cursor     string
}

// SetFatherId Receives the div ID used as a stage for the IDE and puts it to occupy the entire browser area
func (e *cursor) SetFatherId(fatherId string) {
	e.ideStage = factoryBrowser.NewTagDiv().
		Id("cursorFather").
		Import(fatherId)
}

func (e *cursor) SetColor(color color.RGBA) {
	e.vertical.Stroke(color)
	e.horizontal.Stroke(color)
}

func (e *cursor) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

func (e *cursor) GetVisible() (visible string) {
	return e.visible
}

func (e *cursor) Hide() {
	e.visible = "none"

	e.vertical.AddStyle("visibility", "hidden")
	e.horizontal.AddStyle("visibility", "hidden")

	e.ideStage.AddStyle("cursor", e.cursor)
	log.Printf("hide entrou")
}

func (e *cursor) ChangesDirection() {
	if e.visible == "vertical" {
		e.ShowHorizontal()
		return
	}

	e.ShowVertical()
}

func (e *cursor) ShowVertical() {
	e.visible = "vertical"

	e.vertical.AddStyle("visibility", "visible")
	e.horizontal.AddStyle("visibility", "hidden")

	e.cursor = e.ideStage.GetStyle("cursor")
	//e.ideStage.AddStyle("cursor", "none") // todo: descomentar
}

func (e *cursor) ShowHorizontal() {
	e.visible = "horizontal"

	e.horizontal.AddStyle("visibility", "visible")
	e.vertical.AddStyle("visibility", "hidden")

	e.cursor = e.ideStage.GetStyle("cursor")
	//e.ideStage.AddStyle("cursor", "none") // todo: descomentar
}

func (e *cursor) Init() {
	e.document = js.Global().Get("document")
	e.body = e.document.Get("body")

	e.block = factoryBrowser.NewTagDiv().
		Id("cursorBlock").
		//Class(e.classListName).
		AddStyle("position", "absolute").
		AddStyle("top", fmt.Sprintf("%dpx", 0)).
		AddStyle("left", fmt.Sprintf("%dpx", 0)).
		AddStyle("width", fmt.Sprintf("%dpx", e.ideStage.GetClientWidth())).
		AddStyle("height", fmt.Sprintf("%dpx", e.ideStage.GetClientHeight()))
	e.ideStage.Append(e.block)

	e.svg = factoryBrowser.NewTagSvg().
		Id("cursorSvg").
		AddStyle("top", fmt.Sprintf("%dpx", 0)).
		AddStyle("left", fmt.Sprintf("%dpx", 0)).
		AddStyle("width", fmt.Sprintf("%dpx", e.ideStage.GetClientWidth())).
		AddStyle("height", fmt.Sprintf("%dpx", e.ideStage.GetClientHeight()))
	e.block.Append(e.svg)

	e.vertical = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke(factoryColor.NewBlueHalfTransparent()).
		StrokeWidth(2).
		AddStyle("visibility", "hidden")
	e.svg.Append(e.vertical)

	e.horizontal = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke(factoryColor.NewBlueHalfTransparent()).
		StrokeWidth(2).
		AddStyle("visibility", "hidden")
	e.svg.Append(e.horizontal)
}

func (e *cursor) Move(x, y, spaceWidth, spaceHeight int) {
	way := []string{
		fmt.Sprintf("M0,%v", y),
		fmt.Sprintf("L%v,%v", spaceWidth, y),
	}
	e.horizontal.D(way)

	way = []string{
		fmt.Sprintf("M%v,0", x),
		fmt.Sprintf("L%v,%v", x, spaceHeight),
	}
	e.vertical.D(way)
}

func (e *cursor) Update(width, height int) (err error) {
	e.block.AddStyle("width", fmt.Sprintf("%dpx", width))
	e.block.AddStyle("height", fmt.Sprintf("%dpx", height))

	e.svg.AddStyle("width", fmt.Sprintf("%vpx", width))
	e.svg.AddStyle("height", fmt.Sprintf("%vpx", height))

	e.svg.ViewBox([]int{0.0, 0.0, width, height})
	return
}
