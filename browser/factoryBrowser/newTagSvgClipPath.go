package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgClipPath
//
// English:
//
// The <clipPath> SVG element defines a clipping path, to be used by the clip-path property.
//
// A clipping path restricts the region to which paint can be applied. Conceptually, parts of the drawing that lie
// outside of the region bounded by the clipping path are not drawn.
//
// Português:
//
// O elemento SVG <clipPath> define um caminho de recorte, a ser usado pela propriedade clip-path.
//
// Um traçado de recorte restringe a região na qual a tinta pode ser aplicada. Conceitualmente, as partes do desenho
// que estão fora da região delimitada pelo caminho de recorte não são desenhadas.
func NewTagSvgClipPath() (ref *html.TagSvgClipPath) {
	ref = &html.TagSvgClipPath{}
	ref.Init()

	return ref
}
