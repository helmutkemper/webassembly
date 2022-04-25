package factoryColor

import "image/color"

func NewHoneydew() color.RGBA {
	return color.RGBA{R: 0xf0, G: 0xff, B: 0xf0, A: 0xff} // rgb(240, 255, 240)
}
