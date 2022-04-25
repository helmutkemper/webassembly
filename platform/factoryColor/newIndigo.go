package factoryColor

import "image/color"

func NewIndigo() color.RGBA {
	return color.RGBA{R: 0x4b, G: 0x00, B: 0x82, A: 0xff} // rgb(75, 0, 130)
}
