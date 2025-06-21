package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgFeMorphology
//
// English:
//
// The <feMorphology> SVG filter primitive is used to erode or dilate the input image. Its usefulness lies especially
// in fattening or thinning effects.
//
// Português:
//
// A primitiva de filtro SVG <feMorphology> é usada para corroer ou dilatar a imagem de entrada. Sua utilidade reside
// especialmente nos efeitos de engorda ou desbaste.
func NewTagSvgFeMorphology() (ref *html.TagSvgFeMorphology) {
	ref = &html.TagSvgFeMorphology{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
