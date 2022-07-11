package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagImg() (ref *html.TagImg) {
	ref = &html.TagImg{}
	ref.Init()

	return ref
}
