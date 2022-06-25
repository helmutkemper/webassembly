package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

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
func NewTagSvgG(id string) (ref *html.TagSvgG) {
	ref = &html.TagSvgG{}
	ref.Init(id)

	return ref
}
