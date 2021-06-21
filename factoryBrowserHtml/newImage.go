package factoryBrowserHtml

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

func NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad, append bool) interface{} {
	img := html.Html{}
	ret := img.NewImage(parent, propertiesList, waitLoad)

	if append == true {
		img.Append(parent, ret)
	}

	return ret
}
