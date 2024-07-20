package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgUse
//
// English:
//
// The <use> element takes nodes from within the SVG document, and duplicates them somewhere else.
//
// Português:
//
// O elemento <use> pega nós de dentro do documento SVG e os duplica em outro lugar.
func NewTagSvgUse() (ref *html.TagSvgUse) {
	ref = &html.TagSvgUse{}
	ref.Init()

	return ref
}
