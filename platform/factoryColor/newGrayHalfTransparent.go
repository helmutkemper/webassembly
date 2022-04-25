package factoryColor

import "image/color"

func NewGrayHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0x80} // rgb(128, 128, 128)
}
