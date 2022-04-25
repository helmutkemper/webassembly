package factoryColor

import "image/color"

func NewChartreuseTransparent() color.RGBA {
	return color.RGBA{R: 0x7f, G: 0xff, B: 0x00, A: 0x00} // rgb(127, 255, 0)
}
