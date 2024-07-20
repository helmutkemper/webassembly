package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgPolygon
//
// English:
//
// The <polygon> element defines a closed shape consisting of a set of connected straight line segments. The last point
// is connected to the first point.
//
// For open shapes, see the <polyline> element.
//
// Português:
//
// O elemento <polygon> define uma forma fechada que consiste em um conjunto de segmentos de linha reta conectados.
// O último ponto está conectado ao primeiro ponto.
//
// Para formas abertas, consulte o elemento <polyline>.
func NewTagSvgPolygon() (ref *html.TagSvgPolygon) {
	ref = &html.TagSvgPolygon{}
	ref.Init()

	return ref
}
