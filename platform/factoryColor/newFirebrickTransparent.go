package factoryColor

import "image/color"

func NewFirebrickTransparent() color.RGBA {
	return color.RGBA{R: 0xb2, G: 0x22, B: 0x22, A: 0x00} // rgb(178, 34, 34)
}
