package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagCanvas(id string, width, height int) (ref *html.TagCanvas) {
	ref = &html.TagCanvas{}
	ref.CreateElement(html.KTagCanvas, width, height)
	ref.Id(id)

	return ref
}
