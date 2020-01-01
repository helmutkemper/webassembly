package factoryBrowserImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/Html"
)

func NewImage(platform iotmaker_platform_IDraw.IHtml, parent interface{}, propertiesList map[string]interface{}, waitLoad, append bool) Html.Image {
	img := platform.NewImage(parent, propertiesList, waitLoad)

	if append == true {
		platform.Append(parent, img.Get())
	}

	return img
}
