package canvas

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
)

func NewStage(document document.Document, id string, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Stage {
	stage := Stage{}

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(width)
	width = densityCalc.Int()

	densityCalc.Set(height)
	height = densityCalc.Int()

	stage.Canvas = NewCanvasWith2DContext(document.SelfDocument, id, width, height)
	stage.ScratchPad = NewCanvasWith2DContext(document.SelfDocument, id+"ScratchPad", width, height)

	document.AppendChildToDocumentBody(stage.SelfElement)

	return stage
}
