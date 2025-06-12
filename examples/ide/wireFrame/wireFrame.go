package wireFrame

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

type WireFrame struct {
	spiderWay   []coordinate
	bindingType string

	svg       *html.TagSvg
	spiderWeb *html.TagSvgPath
}

func (e *WireFrame) BindingType(binding string) {

}

func (e *WireFrame) AddPoint(x, y int, bindingType string) {
	if e.spiderWay == nil {
		e.spiderWay = make([]coordinate, 0)
	}

	e.spiderWay = append(e.spiderWay, coordinate{x: x, y: y})
}

func (e *WireFrame) GetPointsLength() (length int) {
	return len(e.spiderWay)
}

func (e *WireFrame) Init() (err error) {
	e.svg = factoryBrowser.NewTagSvg()

	e.spiderWeb = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke("red").
		StrokeWidth(3) //.
	//MarkerEnd("url(#spiderWeb)")
	e.svg.Append(e.spiderWeb)

	return
}

func (e *WireFrame) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

func (e *WireFrame) SetWarning(warning bool) {

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

	e.spiderWeb.D(way)

	return
}
