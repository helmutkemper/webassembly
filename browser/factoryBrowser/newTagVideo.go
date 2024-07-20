package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

func NewTagVideo() (ref *html.TagVideo) {
	ref = &html.TagVideo{}
	ref.Init()

	return ref
}
