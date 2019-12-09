package factoryStage

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryCanvas"
)

func NewStage(document document.Document, id string, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) canvas.Stage {
	stage := canvas.Stage{}

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(width)
	width = densityCalc.Int()

	densityCalc.Set(height)
	height = densityCalc.Int()

	stage.Canvas = factoryCanvas.NewCanvasWith2DContext(document.SelfDocument, id, width, height)
	stage.ScratchPad = factoryCanvas.NewCanvasWith2DContext(document.SelfDocument, id+"ScratchPad", width, height)

	document.AppendChildToDocumentBody(stage.SelfElement)

	return stage
}
