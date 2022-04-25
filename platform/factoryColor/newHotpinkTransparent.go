package factoryColor

import "image/color"

func NewHotpinkTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x69, B: 0xb4, A: 0x00} // rgb(255, 105, 180)
}
