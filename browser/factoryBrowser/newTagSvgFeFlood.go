package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgFeFlood
//
// English:
//
// The <feFlood> SVG filter primitive fills the filter subregion with the color and opacity defined by flood-color and
// flood-opacity.
//
// Português:
//
// A primitiva de filtro SVG <feFlood> preenche a sub-região do filtro com a cor e a opacidade definidas por flood-color
// e flood-opacity.
func NewTagSvgFeFlood() (ref *html.TagSvgFeFlood) {
	ref = &html.TagSvgFeFlood{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
