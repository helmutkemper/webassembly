package factoryBrowserStage

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserCanvas"
)

func NewStage(document document.Document, id string, width, height float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) canvas.Stage {
	stage := canvas.Stage{}

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(width)
	width = densityCalc.Float64()

	densityCalc.Set(height)
	height = densityCalc.Float64()

	stage.Canvas = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, id, width, height)
	stage.ScratchPad = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, id+"ScratchPad", width, height)

	document.AppendChildToDocumentBody(stage.SelfElement)

	return stage
}
