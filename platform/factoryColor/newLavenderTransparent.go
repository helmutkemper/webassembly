package factoryColor

import "image/color"

func NewLavenderTransparent() color.RGBA {
	return color.RGBA{R: 0xe6, G: 0xe6, B: 0xfa, A: 0x00} // rgb(230, 230, 250)
}
