package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagLegend
//
// English:
//
//  Create the Legend element.
//
// The <legend> HTML element represents a caption for the content of its parent <fieldset>.
//
// Português:
//
//  Crie o elemento Legenda.
//
// O elemento HTML <legend> representa uma legenda para o conteúdo de seu pai <fieldset>.
func NewTagLegend(id string) (ref *html.TagLegend) {
	ref = &html.TagLegend{}
	ref.CreateElement(html.KTagLegend)
	ref.Id(id)

	return ref
}
