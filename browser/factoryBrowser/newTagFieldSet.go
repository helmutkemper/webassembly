package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

// NewTagFieldSet
//
// English:
//
//	Create the fieldset element.
//
// The <fieldset> HTML element is used to group several controls as well as labels (<label>)
// within a web form.
//
// Português:
//
//	Cria o elemento fieldset.
//
// O elemento HTML <fieldset> é usado para agrupar vários controles, bem como rótulos (<label>)
// dentro de um formulário web.
func NewTagFieldSet() (ref *html.TagFieldset) {
	ref = &html.TagFieldset{}
	ref.CreateElement(html.KTagFieldset)
	ref.Id(mathUtil.GetUID())

	return ref
}
