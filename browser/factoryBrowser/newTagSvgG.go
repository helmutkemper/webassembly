package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgG
//
// English:
//
// The <g> SVG element is a container used to group other SVG elements.
//
// Transformations applied to the <g> element are performed on its child elements, and its attributes are inherited by
// its children. It can also group multiple elements to be referenced later with the <use> element.
//
// Português:
//
// O elemento SVG <g> é um contêiner usado para agrupar outros elementos SVG.
//
// As transformações aplicadas ao elemento <g> são realizadas em seus elementos filhos, e seus atributos são herdados
// por seus filhos. Ele também pode agrupar vários elementos para serem referenciados posteriormente com o elemento
// <use>.
func NewTagSvgG() (ref *html.TagSvgG) {
	ref = &html.TagSvgG{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
