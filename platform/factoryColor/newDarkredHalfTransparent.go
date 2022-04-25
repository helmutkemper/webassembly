package factoryColor

import "image/color"

func NewDarkredHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x8b, G: 0x00, B: 0x00, A: 0x80} // rgb(139, 0, 0)
}
