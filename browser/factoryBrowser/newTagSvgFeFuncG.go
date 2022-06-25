package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeFuncG
//
// English:
//
// The <feFuncG> SVG filter primitive defines the transfer function for the green component of the input graphic of
// its parent <feComponentTransfer> element.
//
// Português:
//
// A primitiva de filtro SVG <feFuncG> define a função de transferência para o componente verde do gráfico de entrada
// de seu elemento pai <feComponentTransfer>.
func NewTagSvgFeFuncG(id string) (ref *html.TagSvgFeFuncG) {
	ref = &html.TagSvgFeFuncG{}
	ref.Init(id)

	return ref
}
