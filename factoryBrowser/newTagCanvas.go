package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

func NewTagCanvas(id string, width, height int) (ref *html.TagCanvas) {
	ref = &html.TagCanvas{}
	ref.CreateElement(html.KTagCanvas, width, height)
	ref.Id(id)

	return ref
}
