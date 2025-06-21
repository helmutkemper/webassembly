package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

func NewTagImg() (ref *html.TagImg) {
	ref = &html.TagImg{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
