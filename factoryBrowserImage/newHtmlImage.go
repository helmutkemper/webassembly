package factoryBrowserImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
)

// fixme:
// fixme: width e heght deven ser passados por parametro e ter density

func NewImage(
	platform iotmaker_platform_IDraw.IHtml,
	parent interface{},
	propertiesList map[string]interface{},
	waitLoad,
	append bool,
) html.Image {
	img := platform.NewImage(parent, propertiesList, waitLoad)

	if append == true {
		platform.Append(parent, img.Get())
	}

	return img
}
