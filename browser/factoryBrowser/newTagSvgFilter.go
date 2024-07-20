package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgFilter
//
// English:
//
// The <filter> SVG element defines a custom filter effect by grouping atomic filter primitives. It is never rendered
// itself, but must be used by the filter attribute on SVG elements, or the filter CSS property for SVG/HTML elements.
//
// Português:
//
// O elemento SVG <filter> define um efeito de filtro personalizado agrupando primitivos de filtro atômico. Ele nunca é
// renderizado, mas deve ser usado pelo atributo filter em elementos SVG ou pela propriedade CSS filter para elementos
// SVGHTML.
func NewTagSvgFilter() (ref *html.TagSvgFilter) {
	ref = &html.TagSvgFilter{}
	ref.Init()

	return ref
}
