package factoryColor

import "image/color"

func NewGreenyellowTransparent() color.RGBA {
	return color.RGBA{R: 0xad, G: 0xff, B: 0x2f, A: 0x00} // rgb(173, 255, 47)
}
