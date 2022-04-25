package factoryColor

import "image/color"

func NewWheatHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xf5, G: 0xde, B: 0xb3, A: 0x80} // rgb(245, 222, 179)
}
