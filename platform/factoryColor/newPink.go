package factoryColor

import "image/color"

func NewPink() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xc0, B: 0xcb, A: 0xff} // rgb(255, 192, 203)
}
