package canvas

import (
	"image/color"
	//"syscall/js"
)

func (el *Canvas) SetPixel(x, y int, pixelColor color.RGBA) {
	pixel := el.SelfContext.Call("createImageData", 1, 1)

	data := pixel.Get("data")

	data.SetIndex(0, pixelColor.R)
	data.SetIndex(1, pixelColor.G)
	data.SetIndex(2, pixelColor.B)
	data.SetIndex(3, pixelColor.A)

	pixel.Set("data", data)

	el.SelfContext.Call("putImageData", pixel, x, y)
}
