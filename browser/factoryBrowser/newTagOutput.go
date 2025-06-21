package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagOutput
//
// English:
//
//	The <output> HTML element is a container element into which a site or app can inject the results
//	of a calculation or the outcome of a user action.
//
// Português:
//
//	O elemento HTML <output> é um elemento de contêiner no qual um site ou aplicativo pode injetar os
//	resultados de um cálculo ou o resultado de uma ação do usuário.
func NewTagOutput() (ref *html.TagOutput) {
	ref = &html.TagOutput{}
	ref.CreateElement(html.KTagMeter)
	ref.Id(utilsMath.GetUID())

	return ref
}
