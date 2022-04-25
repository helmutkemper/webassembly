package factoryColor

import "image/color"

func NewThistleTransparent() color.RGBA {
	return color.RGBA{R: 0xd8, G: 0xbf, B: 0xd8, A: 0x00} // rgb(216, 191, 216)
}
