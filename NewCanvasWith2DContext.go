package iotmaker_platform_webbrowser

func NewCanvasWith2DContext(id string, width, height float64) Canvas {
	el := Canvas{}
	el.InitializeContext2DById(id)

	//el.selfCanvas = el.selfElement

	el.selfCanvas.Set("width", width)
	el.selfCanvas.Set("height", height)
	//el.selfContext = el.selfCanvas.Call("getContext", "2d")

	return el
}
