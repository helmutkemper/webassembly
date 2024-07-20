package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

func NewTagSource() (ref *html.TagSource) {
	ref = &html.TagSource{}
	ref.Init()

	return ref
}
