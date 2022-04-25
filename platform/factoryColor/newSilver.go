package factoryColor

import "image/color"

func NewSilver() color.RGBA {
	return color.RGBA{R: 0xc0, G: 0xc0, B: 0xc0, A: 0xff} // rgb(192, 192, 192)
}
