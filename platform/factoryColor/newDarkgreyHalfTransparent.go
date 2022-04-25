package factoryColor

import "image/color"

func NewDarkgreyHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xa9, G: 0xa9, B: 0xa9, A: 0x80} // rgb(169, 169, 169)
}
