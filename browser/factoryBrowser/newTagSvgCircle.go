package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgCircle
//
// English:
//
// The <circle> SVG element is an SVG basic shape, used to draw circles based on a center point and a radius.
//
// Português:
//
// O elemento SVG <circle> é uma forma básica SVG, usada para desenhar círculos com base em um ponto central e um raio.
func NewTagSvgCircle(id string) (ref *html.TagSvgCircle) {
	ref = &html.TagSvgCircle{}
	ref.Init(id)

	return ref
}
