package factoryColor

import "image/color"

func NewDarkredTransparent() color.RGBA {
	return color.RGBA{R: 0x8b, G: 0x00, B: 0x00, A: 0x00} // rgb(139, 0, 0)
}
