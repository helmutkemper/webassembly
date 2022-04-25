package factoryColor

import "image/color"

func NewSilverHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xc0, G: 0xc0, B: 0xc0, A: 0x80} // rgb(192, 192, 192)
}
