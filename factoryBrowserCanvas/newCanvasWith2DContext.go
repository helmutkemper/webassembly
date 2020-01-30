package factoryBrowserCanvas

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"syscall/js"
)

func NewCanvasWith2DContext(document js.Value, id string, width, height int) canvas.Canvas {
	el := canvas.Canvas{}
	el.SelfElement = document

	el.SelfElement = el.SelfElement.Call("createElement", "canvas")
	el.SelfElement.Set("id", id)

	el.SelfElement.Set("width", width)
	el.SelfElement.Set("height", height)
	el.SelfContext = el.SelfElement.Call("getContext", "2d")

	return el
}
