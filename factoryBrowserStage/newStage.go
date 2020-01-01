package factoryBrowserStage

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserCanvas"
	"github.com/helmutkemper/iotmaker.platform/fps"
)

func NewStage(document document.Document, id string, width, height float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) canvas.Stage {
	stage := canvas.Stage{}

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(width)
	stage.Width = densityCalc.Float64()

	densityCalc.Set(height)
	stage.Height = densityCalc.Float64()

	stage.Canvas = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, id, stage.Width, stage.Height)
	stage.ScratchPad = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, id+"ScratchPad", stage.Width, stage.Height)

	document.HideMousePointer()
	document.AppendChildToDocumentBody(stage.SelfElement)

	fps.AddToRunnerPriorityFunc(stage.Clear)

	return stage
}
