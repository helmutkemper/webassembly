package factoryColor

import "image/color"

func NewMaroonHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x80, G: 0x00, B: 0x00, A: 0x80} // rgb(128, 0, 0)
}
