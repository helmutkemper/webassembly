package canvas

import (
	"image/color"
)

func (el *Canvas) CreateImageData(width, height int, pixelColor color.RGBA) interface{} {
	imageData := el.SelfContext.Call("createImageData", width, height)
	return imageData.Get("data")
}
