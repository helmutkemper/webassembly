package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagCanvas(width, height int) (ref *html.TagCanvas) {
	ref = &html.TagCanvas{}
	ref.Init(width, height)

	return ref
}
