package factoryBrowser

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func NewPoints(list []html.Point) (ref *html.Points) {
	ref = &html.Points{}
	*ref = html.Points(list)

	return
}
