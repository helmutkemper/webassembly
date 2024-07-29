package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

// NewTagSvgCircle
//
// English:
//
// The <circle> SVG element is an SVG basic shape, used to draw circles based on a center point and a radius.
//
// Português:
//
// O elemento SVG <circle> é uma forma básica SVG, usada para desenhar círculos com base em um ponto central e um raio.
func NewTagSvgCircle() (ref *html.TagSvgCircle) {
	ref = &html.TagSvgCircle{}
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}
