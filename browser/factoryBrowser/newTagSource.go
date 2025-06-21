package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

func NewTagSource() (ref *html.TagSource) {
	ref = &html.TagSource{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
