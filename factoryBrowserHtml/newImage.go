package factoryBrowserHtml

import "github.com/helmutkemper/iotmaker.platform.webbrowser/Html"

func NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad, append bool) interface{} {
	img := Html.Html{}
	ret := img.NewImage(parent, propertiesList, waitLoad)

	if append == true {
		img.Append(parent, ret)
	}

	return ret
}
