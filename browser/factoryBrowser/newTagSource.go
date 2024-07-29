package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

func NewTagSource() (ref *html.TagSource) {
	ref = &html.TagSource{}
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}
