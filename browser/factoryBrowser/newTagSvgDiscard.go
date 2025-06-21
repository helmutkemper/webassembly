package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgDiscard
//
// English:
//
// The <discard> SVG element allows authors to specify the time at which particular elements are to be discarded,
// thereby reducing the resources required by an SVG user agent. This is particularly useful to help SVG viewers
// conserve memory while displaying long-running documents.
//
// The <discard> element may occur wherever the <animate> element may.
//
// Português:
//
// The <discard> SVG element allows authors to specify the time at which particular elements are to be discarded,
// thereby reducing the resources required by an SVG user agent. This is particularly useful to help SVG viewers
// conserve memory while displaying long-running documents.
//
// The <discard> element may occur wherever the <animate> element may.
func NewTagSvgDiscard() (ref *html.TagSvgDiscard) {
	ref = &html.TagSvgDiscard{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
