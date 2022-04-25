package factoryColor

import "image/color"

func NewGreenHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0x80} // rgb(0, 128, 0)
}
