package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagSvgRect(id string) (ref *html.TagSvgRect) {
	ref = &html.TagSvgRect{}
	ref.Init(id)

	return ref
}
