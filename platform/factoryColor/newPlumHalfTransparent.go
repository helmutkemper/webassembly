package factoryColor

import "image/color"

func NewPlumHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xdd, G: 0xa0, B: 0xdd, A: 0x80} // rgb(221, 160, 221)
}
