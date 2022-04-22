package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

func NewTagMeter(id string) (ref *html.Meter) {
	ref = &html.Meter{}
	ref.CreateElement(html.KTagMeter)
	ref.SetId(id)

	return ref
}
