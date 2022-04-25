package factoryColor

import "image/color"

func NewSkyblueTransparent() color.RGBA {
	return color.RGBA{R: 0x87, G: 0xce, B: 0xeb, A: 0x00} // rgb(135, 206, 235)
}
