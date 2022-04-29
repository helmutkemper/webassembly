package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browser/html"

// NewTagOutput
//
// English:
//
//  The <output> HTML element is a container element into which a site or app can inject the results
//  of a calculation or the outcome of a user action.
//
// Português:
//
//  O elemento HTML <output> é um elemento de contêiner no qual um site ou aplicativo pode injetar os
//  resultados de um cálculo ou o resultado de uma ação do usuário.
func NewTagOutput(id string) (ref *html.TagOutput) {
	ref = &html.TagOutput{}
	ref.CreateElement(html.KTagMeter)
	ref.Id(id)

	return ref
}
