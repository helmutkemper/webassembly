package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgFeMerge
//
// English:
//
// The <feMerge> SVG element allows filter effects to be applied concurrently instead of sequentially. This is achieved
// by other filters storing their output via the result attribute and then accessing it in a <feMergeNode> child.
//
// Português:
//
// O elemento SVG <feMerge> permite que efeitos de filtro sejam aplicados simultaneamente em vez de sequencialmente.
// Isso é obtido por outros filtros armazenando sua saída por meio do atributo result e acessando-a em um filho
// <feMergeNode>.
func NewTagSvgFeMerge() (ref *html.TagSvgFeMerge) {
	ref = &html.TagSvgFeMerge{}
	ref.Init()

	return ref
}
