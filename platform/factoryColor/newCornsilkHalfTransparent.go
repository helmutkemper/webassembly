package factoryColor

import "image/color"

func NewCornsilkHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xf8, B: 0xdc, A: 0x80} // rgb(255, 248, 220)
}
