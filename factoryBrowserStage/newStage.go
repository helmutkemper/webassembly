package factoryBrowserStage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserCanvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/fps"
)

func NewStage(htmlPlatform iotmaker_platform_IDraw.IHtml, document document.Document, id string, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *canvas.Stage {
	stage := &canvas.Stage{}
	stage.Id = id

	stage.Width = float64(document.GetDocumentWidth())
	stage.Height = float64(document.GetDocumentHeight())

	stage.Canvas = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, stage.Id, stage.Width, stage.Height)
	stage.ScratchPad = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, stage.Id+"ScratchPad", stage.Width, stage.Height)

	document.HideMousePointer()
	document.AppendChildToDocumentBody(stage.SelfElement)

	stage.Canvas.SetWidth(stage.Width)
	stage.Canvas.SetHeight(stage.Height)

	stage.ScratchPad.SetWidth(stage.Width)
	stage.ScratchPad.SetHeight(stage.Height)

	stage.AddToFpsFunc(fps.Set)
	stage.AddToFpsCacheFunc(fps.SetCacheUpdate)

	stage.AddToRunnerFunc(fps.AddToRunner)
	stage.DeleteFromRunnerFunc(fps.DeleteFromRunner)
	stage.AddToCacheRunnerFunc(fps.AddToCacheRunner)
	stage.DeleteFromCacheRunnerFunc(fps.DeleteFromCacheRunner)

	stage.AddToRunnerPriorityFunc(fps.AddToRunnerPriorityFunc)
	stage.DeleteFromRunnerPriorityFunc(fps.DeleteFromRunnerPriorityFunc)
	stage.AddLowLatencyFunc(fps.AddLowLatencyFunc)
	stage.DeleteLowLatencyFunc(fps.DeleteLowLatencyFunc)

	stage.AddWidthPriority(stage.Clear)
	stage.AddWidthLowLatency(func() {
		if document.GetDocumentWidth() != int(stage.Width) || document.GetDocumentHeight() != int(stage.Height) {
			stage.Width = float64(document.GetDocumentWidth())
			stage.Height = float64(document.GetDocumentHeight())

			stage.Canvas.SetWidth(stage.Width)
			stage.Canvas.SetHeight(stage.Height)

			stage.ScratchPad.SetWidth(stage.Width)
			stage.ScratchPad.SetHeight(stage.Height)
		}
	})

	PreLoadCursor(
		document.SelfDocument,
		htmlPlatform,
		&stage.Canvas,
		&stage.ScratchPad,
		KTemplarianPath,
		KTemplarianList,
		density,
		iDensity,
	)

	stage.SetCursorDrawFunc(imageCursor.Draw)
	stage.SetCursor = SetCursor

	stage.SetCursorStageId(stage.AddToStage(imageCursor.Draw))

	return stage
}
