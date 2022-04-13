package factoryBrowserStage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/config"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserCanvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/javascript/canvas"
	webBrowserMouse "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
)

// fixme: stage está estranho

func NewStage(
	htmlPlatform iotmaker_platform_IDraw.IHtml,
	engine engine.IEngine,
	document document.Document,
	id string,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity,
) (
	stage *canvas.Stage,
) {

	stage = &canvas.Stage{}
	stage.Id = id

	stage.Engine = engine
	stage.Engine.Init()

	stage.Width = document.GetDocumentWidth()
	stage.Height = document.GetDocumentHeight()

	stage.Canvas = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, stage.Id, stage.Width, stage.Height)
	stage.ScratchPad = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, stage.Id+"ScratchPad", stage.Width, stage.Height)
	stage.Cache = factoryBrowserCanvas.NewCanvasWith2DContext(document.SelfDocument, stage.Id+"Cache", stage.Width, stage.Height)

	document.MousePointerHide()
	document.AppendToDocument(stage.SelfElement)

	stage.Canvas.SetWidth(stage.Width)
	stage.Canvas.SetHeight(stage.Height)

	stage.ScratchPad.SetWidth(stage.Width)
	stage.ScratchPad.SetHeight(stage.Height)

	stage.Cache.SetWidth(stage.Width)
	stage.Cache.SetHeight(stage.Height)

	stage.AddToSystem(stage.Clear)
	stage.AddToHighLatency(func() {
		if document.GetDocumentWidth() != int(stage.Width) || document.GetDocumentHeight() != int(stage.Height) {
			stage.Width = document.GetDocumentWidth()
			stage.Height = document.GetDocumentHeight()

			stage.Canvas.SetWidth(stage.Width)
			stage.Canvas.SetHeight(stage.Height)

			stage.ScratchPad.SetWidth(stage.Width)
			stage.ScratchPad.SetHeight(stage.Height)

			stage.Cache.SetWidth(stage.Width)
			stage.Cache.SetHeight(stage.Height)
		}
	})

	PreLoadCursor(
		document.SelfDocument,
		stage,
		htmlPlatform,
		&stage.Canvas,
		&stage.ScratchPad,
		config.KTemplarianPath,
		KTemplarianList,
		density,
		iDensity,
	)

	stage.SetCursorDrawFunc(imageCursor.Draw)
	stage.SetCursor = SetCursor

	stage.Engine.CursorAddDrawFunction(imageCursor.Draw)

	// pt_br: Mostra o cursor do mouse sempre que o mesmo entra no documento
	document.AddEventListener(eventMouse.KMouseEnter, webBrowserMouse.SetMouseSimpleEventManager(stage.CursorShow))

	// pt_br: Esconde de cursor do mouse sempre que o mesmo sai de cima do documento
	document.AddEventListener(eventMouse.KMouseOut, webBrowserMouse.SetMouseSimpleEventManager(stage.CursorHide))

	//document.AddEventListener(eventMouse.KMouseMove, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	//document.AddEventListener(eventMouse.KMouseEnter, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseEnter))
	//document.AddEventListener(eventMouse.KMouseOut, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseOut))
	//document.AddEventListener(eventMouse.KMouseUp, webBrowserMouse.SetMouseUpEvent(mouse.ManagerMouseUp))
	//document.AddEventListener(eventMouse.KMouseDown, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseDown))
	//document.AddEventListener(eventMouse.KClick, webBrowserMouse.SetMouseMoveManager(mouse.ManagerClick))
	//document.AddEventListener(eventMouse.KDoubleClick, webBrowserMouse.SetMouseMoveManager(mouse.ManagerDoubleClick))
	//document.AddEventListener(eventMouse.KContextMenu, webBrowserMouse.SetMouseMoveManager(mouse.ManagerContextMenu))

	return
}
