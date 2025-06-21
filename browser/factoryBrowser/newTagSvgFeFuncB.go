package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

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
func NewTagSvgFeFuncB() (ref *html.TagSvgFeFuncB) {
	ref = &html.TagSvgFeFuncB{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
