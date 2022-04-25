package factoryColor

import "image/color"

func NewOrangeHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xa5, B: 0x00, A: 0x80} // rgb(255, 165, 0)
}
