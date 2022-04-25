package factoryColor

import "image/color"

func NewBlueTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0x00} // rgb(0, 0, 255)
}
