package factoryColor

import "image/color"

func NewNavyTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x00, B: 0x80, A: 0x00} // rgb(0, 0, 128)
}
