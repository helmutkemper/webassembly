package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

// NewTagFieldSet
//
// English:
//
//  Create the fieldset element.
//
// The <fieldset> HTML element is used to group several controls as well as labels (<label>)
// within a web form.
//
// Português
//
//  Cria o elemento fieldset.
//
// O elemento HTML <fieldset> é usado para agrupar vários controles, bem como rótulos (<label>)
// dentro de um formulário web.
func NewTagFieldSet(id string) (ref *html.A) {
	ref = &html.A{}
	ref.CreateElement(html.KTagFieldset)
	ref.SetId(id)

	return ref
}
