package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeFuncB
//
// English:
//
// The <feFuncB> SVG filter primitive defines the transfer function for the blue component of the input graphic of its
// parent <feComponentTransfer> element.
//
// Português:
//
// The <feFuncB> SVG filter primitive defines the transfer function for the blue component of the input graphic of its
// parent <feComponentTransfer> element.
func NewTagSvgFeFuncB(id string) (ref *html.TagSvgFeFuncB) {
	ref = &html.TagSvgFeFuncB{}
	ref.Init(id)

	return ref
}
