package iotmaker_platform_webbrowser

func NewCanvasWith2DContext(id string, width, height int) Canvas {
	el := Canvas{}
	el.SelfElement = el.Element.SelfElement

	el.InitializeContext2DById(id)

	el.SelfElement.Set("width", width)
	el.SelfElement.Set("height", height)
	el.SelfContext = el.SelfElement.Call("getContext", "2d")

	return el
}
