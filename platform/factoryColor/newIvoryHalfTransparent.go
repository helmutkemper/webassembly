package factoryColor

import "image/color"

func NewIvoryHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xff, B: 0xf0, A: 0x80} // rgb(255, 255, 240)
}
