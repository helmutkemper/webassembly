package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagSvg(id string) (ref *html.TagSvg) {
	ref = &html.TagSvg{}
	ref.Init(id)

	return ref
}
