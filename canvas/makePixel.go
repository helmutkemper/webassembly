package canvas

import (
	"image/color"
	//"syscall/js"
)

func (el *Canvas) MakePixel(pixelColor color.RGBA) interface{} {
	pixel := el.SelfContext.Call("createImageData", 1, 1)

	data := pixel.Get("data")

	data.SetIndex(0, pixelColor.R)
	data.SetIndex(1, pixelColor.G)
	data.SetIndex(2, pixelColor.B)
	data.SetIndex(3, pixelColor.A)

	pixel.Set("data", data)

	return pixel
}
