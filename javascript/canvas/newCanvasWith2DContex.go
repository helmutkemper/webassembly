package canvas

import (
	"syscall/js"
)

func (el *Canvas) NewCanvasWith2DContext(document interface{}, id string, width, height int) (canvas *Canvas) {
	el.SelfElement = document.(js.Value)

	el.SelfElement = el.SelfElement.Call("createElement", "canvas")
	el.SelfElement.Set("id", id)
	el.SelfElement.Set("width", width)
	el.SelfElement.Set("height", height)
	el.SelfContext = el.SelfElement.Call("getContext", "2d")

	return el
}
