package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

// NewTagA
//
// English:
//
//  Create the Anchor element.
//
// The <a> HTML element (or anchor element), with its href attribute, creates a hyperlink to web
// pages, files, email addresses, locations in the same page, or anything else a URL can address.
//
// Content within each <a> should indicate the link's destination. If the href attribute is present,
// pressing the enter key while focused on the <a> element will activate it.
//
// Português
//
//  Cria o elemento Âncora.
//
// O elemento HTML <a> (ou elemento âncora), com seu atributo href, cria um hiperlink para páginas
// da web, arquivos, endereços de e-mail, locais na mesma página ou qualquer outra coisa que um URL
// possa endereçar.
//
// O conteúdo de cada <a> deve indicar o destino do link. Se o atributo href estiver presente,
// pressionar a tecla enter enquanto estiver focado no elemento <a> irá ativá-lo.
func NewTagA(id string) (ref *html.TagA) {
	ref = &html.TagA{}
	ref.CreateElement(html.KTagA)
	ref.Id(id)

	return ref
}
