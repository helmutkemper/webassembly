package factoryColor

import "image/color"

func NewPeruHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xcd, G: 0x85, B: 0x3f, A: 0x80} // rgb(205, 133, 63)
}
