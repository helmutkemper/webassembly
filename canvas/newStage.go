package canvas

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
)

func NewStage(document document.Document, id string, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Stage {
	stage := Stage{}

	densityWidth := iDensity
	densityWidth.Set(width)
	densityWidth.SetDensityFactor(density)

	densityHeight := iDensity
	densityHeight.Set(height)
	densityHeight.SetDensityFactor(density)

	stage.Canvas = NewCanvasWith2DContext(document.SelfDocument, id, densityWidth.Int(), densityHeight.Int())
	stage.ScratchPad = NewCanvasWith2DContext(document.SelfDocument, id+"ScratchPad", densityWidth.Int(), densityHeight.Int())

	document.AppendChildToDocumentBody(stage.SelfElement)

	return stage
}
