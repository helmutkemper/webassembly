package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

func NewTagTextArea() (ref *html.TagTextArea) {
	ref = &html.TagTextArea{}
	ref.CreateElement(html.KTagTextarea)

	return ref
}
