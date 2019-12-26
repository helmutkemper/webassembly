package factoryBrowserImage

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

func NewImage(platform iotmaker_platform_IDraw.IHtml, parent interface{}, propertiesList map[string]interface{}, waitLoad, append bool) interface{} {
	img := platform.NewImage(parent, propertiesList, waitLoad)

	if append == true {
		platform.Append(parent, img)
	}

	return img
}
