package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

func NewTagHr() (ref *html.TagHr) {
	ref = &html.TagHr{}
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}
