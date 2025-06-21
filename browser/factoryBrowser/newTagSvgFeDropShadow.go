package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgFeDropShadow
//
// English:
//
// The SVG <feDropShadow> filter primitive creates a drop shadow of the input image. It can only be used inside a
// <filter> element.
//
//	Notes:
//	  * The drop shadow color and opacity can be changed by using the flood-color and flood-opacity presentation
//	    attributes.
//
// Português:
//
// A primitiva de filtro SVG <feDropShadow> cria uma sombra projetada da imagem de entrada. Ele só pode ser usado
// dentro de um elemento <filter>.
//
//	Notas:
//	  * A cor e a opacidade da sombra projetada podem ser alteradas usando os atributos de apresentação de cor de
//	    inundação e opacidade de inundação.
func NewTagSvgFeDropShadow() (ref *html.TagSvgFeDropShadow) {
	ref = &html.TagSvgFeDropShadow{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
