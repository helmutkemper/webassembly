package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagTextArea(id string) (ref *html.TagTextArea) {
	ref = &html.TagTextArea{}
	ref.CreateElement(html.KTagMeter)
	ref.Id(id)

	return ref
}
