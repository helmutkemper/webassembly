package factoryColor

import "image/color"

func NewLinenTransparent() color.RGBA {
	return color.RGBA{R: 0xfa, G: 0xf0, B: 0xe6, A: 0x00} // rgb(250, 240, 230)
}
