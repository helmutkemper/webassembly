package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

// NewTagLabel
//
// English:
//
//  Create the Label element.
//
// The <label> HTML element represents a caption for an item in a user interface.
//
// Português:
//
//  Cria o elemento Label.
//
// O elemento HTML <label> representa uma legenda para um item em uma interface do usuário.
func NewTagLabel(id string) (ref *html.TagLabel) {
	ref = &html.TagLabel{}
	ref.CreateElement(html.KTagLabel)
	ref.Id(id)

	return ref
}
