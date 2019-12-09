package factoryBrowserHtml

import "github.com/helmutkemper/iotmaker.platform.webbrowser/Html"

func NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad bool) interface{} {
	img := Html.Html{}
	return img.NewImage(parent, propertiesList, waitLoad)
}
