package canvas

import (
	"image/color"
)

func (el *Canvas) CreateImageData(width, height float64, pixelColor color.RGBA) interface{} {
	imageData := el.SelfContext.Call("createImageData", width, height)
	return imageData.Get("data")
}
