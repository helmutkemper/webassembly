package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

func NewTagVideo() (ref *html.TagVideo) {
	ref = &html.TagVideo{}
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}
