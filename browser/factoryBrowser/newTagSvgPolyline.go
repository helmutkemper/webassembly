package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgPolyline
//
// English:
//
// The <polyline> SVG element is an SVG basic shape that creates straight lines connecting several points. Typically a
// polyline is used to create open shapes as the last point doesn't have to be connected to the first point.
//
// For closed shapes see the <polygon> element.
//
// Português:
//
// O elemento SVG <polyline> é uma forma básica SVG que cria linhas retas conectando vários pontos. Normalmente, uma
// polilinha é usada para criar formas abertas, pois o último ponto não precisa ser conectado ao primeiro ponto.
//
// Para formas fechadas veja o elemento <polygon>.
//
// Para formas abertas, consulte o elemento <polyline>.
func NewTagSvgPolyline() (ref *html.TagSvgPolyline) {
	ref = &html.TagSvgPolyline{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
