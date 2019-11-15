package iotmaker_platform_webbrowser

func NewCanvasWith2DContext(id string, width, height float64) Canvas {
	el := Canvas{}
	el.Element.Initialize()
	el.Element.NewCanvas(id)

	el.selfCanvas = el.selfElement

	el.selfCanvas.Set("width", width)
	el.selfCanvas.Set("height", height)
	el.selfCanvas = el.selfCanvas.Call("getContext", "2d")

	return el
}
