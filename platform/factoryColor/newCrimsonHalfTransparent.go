package factoryColor

import "image/color"

func NewCrimsonHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xdc, G: 0x14, B: 0x3c, A: 0x80} // rgb(220, 20, 60)
}
