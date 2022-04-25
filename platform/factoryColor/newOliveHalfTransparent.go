package factoryColor

import "image/color"

func NewOliveHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x80, G: 0x80, B: 0x00, A: 0x80} // rgb(128, 128, 0)
}
