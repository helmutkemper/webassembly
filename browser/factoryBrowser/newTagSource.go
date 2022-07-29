package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagSource() (ref *html.TagSource) {
	ref = &html.TagSource{}
	ref.Init()

	return ref
}
