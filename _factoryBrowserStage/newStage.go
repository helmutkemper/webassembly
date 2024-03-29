package factoryBrowserStage

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"github.com/helmutkemper/iotmaker.webassembly/config"
	"github.com/helmutkemper/iotmaker.webassembly/factoryBrowserCanvas"
	"github.com/helmutkemper/iotmaker.webassembly/globalDocument"
	"github.com/helmutkemper/iotmaker.webassembly/javascript/canvas"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.webassembly/platform/IDraw"
)

// fixme: stage está estranho

func NewStage(
	htmlPlatform iotmaker_platform_IDraw.IHtml,
	engine engine.IEngine,
	document globalDocument.Document,
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

	stage.Canvas = factoryBrowserCanvas.NewCanvasWith2DContext(document.Get(), stage.Id, stage.Width, stage.Height)
	stage.ScratchPad = factoryBrowserCanvas.NewCanvasWith2DContext(document.Get(), stage.Id+"ScratchPad", stage.Width, stage.Height)
	stage.Cache = factoryBrowserCanvas.NewCanvasWith2DContext(document.Get(), stage.Id+"Cache", stage.Width, stage.Height)

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
		document.Get(),
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
	document.AddEventListener(mouse.KEventMouseEnter, mouse.SetMouseSimpleEventManager(stage.CursorShow))

	// pt_br: Esconde de cursor do mouse sempre que o mesmo sai de cima do documento
	document.AddEventListener(mouse.KEventMouseOut, mouse.SetMouseSimpleEventManager(stage.CursorHide))

	//document.AddEventListener(browserMouse.KEventMouseMove, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	//document.AddEventListener(browserMouse.KEventMouseEnter, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseEnter))
	//document.AddEventListener(browserMouse.KEventMouseOut, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseOut))
	//document.AddEventListener(browserMouse.KEventMouseUp, webBrowserMouse.SetMouseUpEvent(mouse.ManagerMouseUp))
	//document.AddEventListener(browserMouse.KEventMouseDown, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseDown))
	//document.AddEventListener(browserMouse.KEventClick, webBrowserMouse.SetMouseMoveManager(mouse.ManagerClick))
	//document.AddEventListener(browserMouse.KEventDoubleClick, webBrowserMouse.SetMouseMoveManager(mouse.ManagerDoubleClick))
	//document.AddEventListener(browserMouse.KEventContextMenu, webBrowserMouse.SetMouseMoveManager(mouse.ManagerContextMenu))

	return
}
