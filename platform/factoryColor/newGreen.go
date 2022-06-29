package factoryColor

import "image/color"

func NewGreen() color.RGBA {
	return color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xff} // rgb(0, 128, 0)
}
