package factoryBrowserImage

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.webassembly/html"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.webassembly/platform/IDraw"
)

// fixme:

func NewImage(
	width int,
	height int,
	propertiesList map[string]interface{},
	waitLoad,
	append bool,
) html.Image {

	var platform iotmaker_platform_IDraw.IHtml = global.Global.Html
	var parent interface{} = global.Global.Document.selfDocument

	densityCalc := global.Global.DensityManager
	densityCalc.SetDensityFactor(global.Global.Density)

	densityCalc.SetInt(width)
	width = densityCalc.Int()

	densityCalc.SetInt(height)
	height = densityCalc.Int()

	propertiesList["width"] = width
	propertiesList["height"] = height

	img := platform.NewImage(parent, propertiesList, waitLoad)

	if append == true {
		platform.Append(parent, img.Get())
	}

	return img
}
