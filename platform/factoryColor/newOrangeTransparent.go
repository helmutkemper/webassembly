package factoryColor

import "image/color"

func NewOrangeTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xa5, B: 0x00, A: 0x00} // rgb(255, 165, 0)
}
