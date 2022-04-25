package factoryColor

import "image/color"

func NewSpringgreenHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0xff, B: 0x7f, A: 0x80} // rgb(0, 255, 127)
}
