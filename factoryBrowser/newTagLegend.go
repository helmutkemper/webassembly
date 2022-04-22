package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

func NewTagLegend(id string) (ref *html.Legend) {
	ref = &html.Legend{}
	ref.CreateElement(html.KTagLegend)
	ref.SetId(id)

	return ref
}
