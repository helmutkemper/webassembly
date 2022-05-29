package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagSvgImage(id string) (ref *html.TagSvgImage) {
	ref = &html.TagSvgImage{}
	ref.Init(id)

	return ref
}
