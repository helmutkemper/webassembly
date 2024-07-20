package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagDataList
//
// English:
//
//	Create the elemento datalist.
//
// The <datalist> HTML element contains a set of <option> elements that represent the permissible
// or recommended options available to choose from within other controls.
//
// Português:
//
//	Cria o elemento datalist
//
// O elemento HTML <datalist> contém um conjunto de elementos <option> que representam as opções
// permitidas ou recomendadas disponíveis para escolha em outros controles.
//
// Example:
//
//	factoryBrowser.NewTagDataList("test_A").
//	  SetNewOption("test_A_a", "label a", "value_a", true, false).
//	  SetNewOption("test_A_b", "label b", "value_b", false, false).
//	  SetNewOption("test_A_c", "label c", "value_c", false, false).
//	  SetNewOption("test_A_d", "label d", "value_d", false, true).
//	  AppendById("stage")
func NewTagDataList(id string) (ref *html.TagDataList) {
	ref = &html.TagDataList{}
	ref.CreateElement(html.KTagDatalist)
	ref.Id(id)

	return ref
}
