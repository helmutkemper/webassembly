package factoryColor

import "image/color"

func NewChocolateHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xd2, G: 0x69, B: 0x1e, A: 0x80} // rgb(210, 105, 30)
}
