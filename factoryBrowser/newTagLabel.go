package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

func NewTagLabel(id string) (ref *html.Label) {
	ref = &html.Label{}
	ref.CreateElement(html.KTagLabel)
	ref.SetId(id)

	return ref
}
