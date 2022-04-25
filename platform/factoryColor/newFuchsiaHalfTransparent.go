package factoryColor

import "image/color"

func NewFuchsiaHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0x80} // rgb(255, 0, 255)
}
