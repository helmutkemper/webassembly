package factoryColor

import "image/color"

func NewSalmonHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xfa, G: 0x80, B: 0x72, A: 0x80} // rgb(250, 128, 114)
}
