package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

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
func NewTagSvgFeFlood(id string) (ref *html.TagSvgFeFlood) {
	ref = &html.TagSvgFeFlood{}
	ref.Init(id)

	return ref
}
