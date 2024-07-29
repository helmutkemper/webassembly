package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

// NewTagSvgFeImage
//
// English:
//
// The <feImage> SVG filter primitive fetches image data from an external source and provides the pixel data as output
// (meaning if the external source is an SVG image, it is rasterized.)
//
// Português:
//
// A primitiva de filtro SVG <feImage> busca dados de imagem de uma fonte externa e fornece os dados de pixel como saída
// (ou seja, se a fonte externa for uma imagem SVG, ela será rasterizada).
func NewTagSvgFeImage() (ref *html.TagSvgFeImage) {
	ref = &html.TagSvgFeImage{}
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}
