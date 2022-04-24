package factoryBrowserDocument

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/globalDocument"
)

func NewDocument() globalDocument.Document {
	el := globalDocument.Document{}
	el.Init()

	el.AddEventListener(browserMouse.KEventMouseMove, browserMouse.SetMouseMoveEvent())
	el.AddEventListener(browserMouse.KEventClick, browserMouse.SetMouseClickEvent())
	el.AddEventListener(browserMouse.KEventDoubleClick, browserMouse.SetMouseDoubleClickEvent())
	el.AddEventListener(browserMouse.KEventMouseDown, browserMouse.SetMouseDownEvent())
	el.AddEventListener(browserMouse.KEventMouseUp, browserMouse.SetMouseUpEvent())

	return el
}
