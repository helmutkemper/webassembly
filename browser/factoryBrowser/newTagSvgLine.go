package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgLine
//
// English:
//
// The <line> element is an SVG basic shape used to create a line connecting two points.
//
// Português:
//
// O elemento <line> é uma forma básica SVG usada para criar uma linha conectando dois pontos.
func NewTagSvgLine() (ref *html.TagSvgLine) {
	ref = &html.TagSvgLine{}
	ref.Init()

	return ref
}
