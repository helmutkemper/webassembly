package iotmaker_platform_webbrowser

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"

func NewStage(id string, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Stage {
	stage := Stage{}

	densityWidth := iDensity
	densityWidth.Set(width)
	densityWidth.SetDensityFactor(density)

	densityHeight := iDensity
	densityHeight.Set(height)
	densityHeight.SetDensityFactor(density)

	stage.Canvas = NewCanvasWith2DContext(id, densityWidth.Int(), densityHeight.Int())

	stage.Canvas.AppendElementToDocumentBody()

	return stage
}
