package factoryColor

import "image/color"

func NewTealTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x80, B: 0x80, A: 0x00} // rgb(0, 128, 128)
}
