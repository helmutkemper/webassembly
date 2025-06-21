package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

func NewTagHr() (ref *html.TagHr) {
	ref = &html.TagHr{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
