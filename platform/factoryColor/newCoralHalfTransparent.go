package factoryColor

import "image/color"

func NewCoralHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x7f, B: 0x50, A: 0x80} // rgb(255, 127, 80)
}
