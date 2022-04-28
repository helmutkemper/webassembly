package factoryBrowserDocument

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/globalDocument"
)

func NewDocument() globalDocument.Document {
	el := globalDocument.Document{}
	el.Init()

	el.AddEventListener(mouse.KEventMouseMove, mouse.SetMouseMoveEvent())
	el.AddEventListener(mouse.KEventClick, mouse.SetMouseClickEvent())
	el.AddEventListener(mouse.KEventDoubleClick, mouse.SetMouseDoubleClickEvent())
	el.AddEventListener(mouse.KEventMouseDown, mouse.SetMouseDownEvent())
	el.AddEventListener(mouse.KEventMouseUp, mouse.SetMouseUpEvent())

	return el
}
