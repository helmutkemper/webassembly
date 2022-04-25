package factoryColor

import "image/color"

func NewPurpleHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x80, G: 0x00, B: 0x80, A: 0x80} // rgb(128, 0, 128)
}
