package factoryColor

import "image/color"

func NewPalegreenHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x98, G: 0xfb, B: 0x98, A: 0x80} // rgb(152, 251, 152)
}
