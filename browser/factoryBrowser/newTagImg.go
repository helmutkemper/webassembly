package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

func NewTagImg() (ref *html.TagImg) {
	ref = &html.TagImg{}
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}
