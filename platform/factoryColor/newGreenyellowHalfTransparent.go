package factoryColor

import "image/color"

func NewGreenyellowHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xad, G: 0xff, B: 0x2f, A: 0x80} // rgb(173, 255, 47)
}
