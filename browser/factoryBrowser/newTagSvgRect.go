package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgRect
//
// English:
//
// The <rect> element is a basic SVG shape that draws rectangles, defined by their position, width, and height.
// The rectangles may have their corners rounded.
//
// Português:
//
// O elemento <rect> é uma forma SVG básica que desenha retângulos, definidos por sua posição, largura e altura.
// Os retângulos podem ter seus cantos arredondados.
func NewTagSvgRect(id string) (ref *html.TagSvgRect) {
	ref = &html.TagSvgRect{}
	ref.Init(id)

	return ref
}
