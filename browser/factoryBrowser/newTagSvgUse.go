package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgUse
//
// English:
//
// The <use> element takes nodes from within the SVG document, and duplicates them somewhere else.
//
// Português:
//
// O elemento <use> pega nós de dentro do documento SVG e os duplica em outro lugar.
func NewTagSvgUse(id string) (ref *html.TagSvgUse) {
	ref = &html.TagSvgUse{}
	ref.Init(id)

	return ref
}
