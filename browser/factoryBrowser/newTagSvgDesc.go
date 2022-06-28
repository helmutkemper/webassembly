package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgDesc
//
// English:
//
// The <desc> element provides an accessible, long-text description of any SVG container element or graphics element.
//
// Text in a <desc> element is not rendered as part of the graphic. If the element can be described by visible text,
// it is possible to reference that text with the aria-describedby attribute. If aria-describedby is used, it will take
// precedence over <desc>.
//
// The hidden text of a <desc> element can also be concatenated with the visible text of other elements using multiple
// IDs in an aria-describedby value. In that case, the <desc> element must provide an ID for reference.
//
// Português:
//
// O elemento <desc> fornece uma descrição de texto longo e acessível de qualquer elemento de contêiner SVG ou
// elemento gráfico.
//
// O texto em um elemento <desc> não é renderizado como parte do gráfico. Se o elemento puder ser descrito por texto
// visível, é possível fazer referência a esse texto com o atributo aria-describedby. Se aria-describedby for usado,
// terá precedência sobre <desc>.
//
// O texto oculto de um elemento <desc> também pode ser concatenado com o texto visível de outros elementos usando
// vários IDs em um valor descrito por aria. Nesse caso, o elemento <desc> deve fornecer um ID para referência.
func NewTagSvgDesc() (ref *html.TagSvgDesc) {
	ref = &html.TagSvgDesc{}
	ref.Init()

	return ref
}
