package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

func NewTagVideo() (ref *html.TagVideo) {
	ref = &html.TagVideo{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
