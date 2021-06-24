package factoryBrowserDocument

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventMouse"
	webBrowserMouse "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/mouse"
)

func NewDocument() document.Document {
	el := document.Document{}
	el.Init()

	el.AddEventListener(eventMouse.KMouseMove, webBrowserMouse.SetMouseMoveEvent())
	el.AddEventListener(eventMouse.KClick, webBrowserMouse.SetMouseClickEvent())
	el.AddEventListener(eventMouse.KDoubleClick, webBrowserMouse.SetMouseDoubleClickEvent())
	el.AddEventListener(eventMouse.KMouseDown, webBrowserMouse.SetMouseDownEvent())
	el.AddEventListener(eventMouse.KMouseUp, webBrowserMouse.SetMouseUpEvent())

	return el
}
