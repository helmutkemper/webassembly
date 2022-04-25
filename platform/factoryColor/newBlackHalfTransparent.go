package factoryColor

import "image/color"

func NewBlackHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x80} // rgb(0, 0, 0)
}
