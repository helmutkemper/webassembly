package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

func NewCanvasWith2DContext(width, height float64) Canvas {
	el := Canvas{}

	el.Element = NewElementById("mycanvas")
	//el.selfCanvas = el.selfDocument.Call("createElement", "canvas")
	el.selfElement.Set("width", width)
	el.selfElement.Set("height", height)
	el.selfContext = el.selfElement.Call("getContext", "2d")

	el.selfContext.Set("globalAlpha", 1.0)
	el.selfContext.Set("strokeStyle", "orange")

	el.selfContext.Call("beginPath")
	el.selfContext.Call("moveTo", 0.0, 0.0)
	el.selfContext.Call("lineTo", 50.0, 50.0)
	el.selfContext.Call("stroke")

	return el
}
