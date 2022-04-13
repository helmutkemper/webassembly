package factoryBrowserDocument

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/globalDocument"
)

func NewDocument() globalDocument.Document {
	el := globalDocument.Document{}
	el.Init()

	el.AddEventListener(eventMouse.KMouseMove, browserMouse.SetMouseMoveEvent())
	el.AddEventListener(eventMouse.KClick, browserMouse.SetMouseClickEvent())
	el.AddEventListener(eventMouse.KDoubleClick, browserMouse.SetMouseDoubleClickEvent())
	el.AddEventListener(eventMouse.KMouseDown, browserMouse.SetMouseDownEvent())
	el.AddEventListener(eventMouse.KMouseUp, browserMouse.SetMouseUpEvent())

	return el
}
