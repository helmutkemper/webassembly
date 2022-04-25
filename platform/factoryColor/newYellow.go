package factoryColor

import "image/color"

func NewYellow() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff} // rgb(255, 255, 0)
}
