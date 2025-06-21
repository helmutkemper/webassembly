package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgStop
//
// English:
//
// The SVG <stop> element defines a color and its position to use on a gradient. This element is always a child of a
// <linearGradient> or <radialGradient> element.
//
// Português:
//
// O elemento SVG <stop> define uma cor e sua posição para usar em um gradiente. Este elemento é sempre um filho de um
// elemento <linearGradient> ou <radialGradient>.
func NewTagSvgStop() (ref *html.TagSvgStop) {
	ref = &html.TagSvgStop{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
