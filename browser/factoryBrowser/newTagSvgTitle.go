package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgTitle
//
// English:
//
// The <title> element provides an accessible, short-text description of any SVG container element or graphics element.
//
// Text in a <title> element is not rendered as part of the graphic, but browsers usually display it as a tooltip.
// If an element can be described by visible text, it is recommended to reference that text with an aria-labelledby
// attribute rather than using the <title> element.
//
//	Notes:
//	  * For backward compatibility with SVG 1.1, <title> elements should be the first child element of their parent.
//
// Português:
//
// O elemento <title> fornece uma descrição de texto curto acessível de qualquer elemento de contêiner SVG ou elemento
// gráfico.
//
// O texto em um elemento <title> não é renderizado como parte do gráfico, mas os navegadores geralmente o exibem como
// uma dica de ferramenta. Se um elemento puder ser descrito por texto visível, é recomendável fazer referência a esse
// texto com um atributo aria-labelledby em vez de usar o elemento <title>.
//
//	Notas:
//	  * Para compatibilidade com versões anteriores com SVG 1.1, os elementos <title> devem ser o primeiro elemento
//	    filho de seu pai.
func NewTagSvgTitle() (ref *html.TagSvgTitle) {
	ref = &html.TagSvgTitle{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
