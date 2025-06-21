package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvg
//
// English:
//
// The svg element is a container that defines a new coordinate system and viewport. It is used as the outermost element
// of SVG documents, but it can also be used to embed an SVG fragment inside an SVG or HTML document.
//
//	Notes:
//	  * The xmlns attribute is only required on the outermost svg element of SVG documents. It is unnecessary for inner
//	    svg elements or inside HTML documents.
//
// Português:
//
// O elemento svg é um contêiner que define um novo sistema de coordenadas e viewport. Ele é usado como o elemento mais
// externo dos documentos SVG, mas também pode ser usado para incorporar um fragmento SVG dentro de um documento SVG
// ou HTML.
//
//	Notas:
//	  * O atributo xmlns só é necessário no elemento svg mais externo dos documentos SVG. É desnecessário para
//	    elementos svg internos ou dentro de documentos HTML.
func NewTagSvg() (ref *html.TagSvg) {
	ref = &html.TagSvg{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
