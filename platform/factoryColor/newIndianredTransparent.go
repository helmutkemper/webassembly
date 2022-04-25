package factoryColor

import "image/color"

func NewIndianredTransparent() color.RGBA {
	return color.RGBA{R: 0xcd, G: 0x5c, B: 0x5c, A: 0x00} // rgb(205, 92, 92)
}
