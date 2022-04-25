package factoryColor

import "image/color"

func NewBlueHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0x80} // rgb(0, 0, 255)
}
