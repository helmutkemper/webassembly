package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgTSpan
//
// English:
//
// The SVG <tspan> element defines a subtext within a <text> element or another <tspan> element. It allows for
// adjustment of the style and/or position of that subtext as needed.
//
// Português:
//
// O elemento SVG <tspan> define um subtexto dentro de um elemento <text> ou outro elemento <tspan>. Permite o ajuste
// do estilo eou posição desse subtexto conforme necessário.
func NewTagSvgTSpan() (ref *html.TagSvgTSpan) {
	ref = &html.TagSvgTSpan{}
	ref.Init()

	return ref
}
