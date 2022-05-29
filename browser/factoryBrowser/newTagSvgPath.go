package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagSvgPath(id string) (ref *html.TagSvgPath) {
	ref = &html.TagSvgPath{}
	ref.Init(id)

	return ref
}
