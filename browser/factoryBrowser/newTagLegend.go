package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagLegend
//
// English:
//
//	Create the Legend element.
//
// The <legend> HTML element represents a caption for the content of its parent <fieldset>.
//
// Português:
//
//	Crie o elemento Legenda.
//
// O elemento HTML <legend> representa uma legenda para o conteúdo de seu pai <fieldset>.
func NewTagLegend() (ref *html.TagLegend) {
	ref = &html.TagLegend{}
	ref.CreateElement(html.KTagLegend)
	ref.Id(utilsMath.GetUID())

	return ref
}
