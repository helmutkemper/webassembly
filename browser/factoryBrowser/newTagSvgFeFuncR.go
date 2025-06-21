package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgFeFuncR
//
// English:
//
// The <feFuncR> SVG filter primitive defines the transfer function for the red component of the input graphic of its
// parent <feComponentTransfer> element.
//
// Português:
//
// A primitiva de filtro SVG <feFuncR> define a função de transferência para o componente vermelho do gráfico de
// entrada de seu elemento pai <feComponentTransfer>.
func NewTagSvgFeFuncR() (ref *html.TagSvgFeFuncR) {
	ref = &html.TagSvgFeFuncR{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
