package factoryColor

import "image/color"

func NewGhostwhiteHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xf8, G: 0xf8, B: 0xff, A: 0x80} // rgb(248, 248, 255)
}
