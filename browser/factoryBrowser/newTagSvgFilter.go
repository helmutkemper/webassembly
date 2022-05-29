package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagSvgFilter(id string) (ref *html.TagSvgFilter) {
	ref = &html.TagSvgFilter{}
	ref.Init(id)

	return ref
}
