package iotmaker_platform_webbrowser

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
)

func NewCanvasWith2DContext(id string, width, height iotmaker_types.Pixel) Canvas {
	el := Canvas{}
	el.InitializeContext2DById(id)

	//el.selfCanvas = el.SelfElement

	el.SelfElement.Set("width", width)
	el.SelfElement.Set("height", height)
	//el.SelfContext = el.selfCanvas.Call("getContext", "2d")

	return el
}
