package factoryColor

import "image/color"

func NewLavenderHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xe6, G: 0xe6, B: 0xfa, A: 0x80} // rgb(230, 230, 250)
}
