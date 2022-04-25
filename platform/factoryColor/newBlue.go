package factoryColor

import "image/color"

func NewBlue() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff} // rgb(0, 0, 255)
}
