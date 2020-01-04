package factoryBrowserStage

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserCanvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/fps"
)

func NewStage(document document.Document, id string) *canvas.Stage {
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

	fps.AddToRunnerPriorityFunc(stage.Clear)
	fps.AddLowLatencyFunc(func() {
		if document.GetDocumentWidth() != int(stage.Width) || document.GetDocumentHeight() != int(stage.Height) {
			stage.Width = float64(document.GetDocumentWidth())
			stage.Height = float64(document.GetDocumentHeight())

			stage.Canvas.SetWidth(stage.Width)
			stage.Canvas.SetHeight(stage.Height)

			stage.ScratchPad.SetWidth(stage.Width)
			stage.ScratchPad.SetHeight(stage.Height)
		}
	})

	Html.PreLoadCursor(document.SelfDocument, Html.KTemplarianPath, Html.KTemplarianList)

	return stage
}
