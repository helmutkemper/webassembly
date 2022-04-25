package factoryColor

import "image/color"

func NewRed() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff} // rgb(255, 0, 0)
}
