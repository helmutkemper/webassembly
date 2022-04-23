package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

func NewTagImage(id, src string, width, height int, waitLoad bool) (ref *html.TagImage) {
	ref = &html.TagImage{}

	ref.CreateElement(html.KTagImg, src, width, height, waitLoad)
	ref.Id(id)

	return ref
}
