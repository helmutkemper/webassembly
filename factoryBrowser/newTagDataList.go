package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

// NewTagDataList
//
// English:
//
//  Create the elemento datalist.
//
// The <datalist> HTML element contains a set of <option> elements that represent the permissible
// or recommended options available to choose from within other controls.
//
// Português
//
//  Cria o elemento datalist
//
// O elemento HTML <datalist> contém um conjunto de elementos <option> que representam as opções
// permitidas ou recomendadas disponíveis para escolha em outros controles.
func NewTagDataList(id string) (ref *html.DataList) {
	ref = &html.DataList{}
	ref.CreateElement(html.KTagDatalist)
	ref.SetId(id)

	return ref
}
