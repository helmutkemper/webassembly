package factoryColor

import "image/color"

func NewYellowHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0x80} // rgb(255, 255, 0)
}
