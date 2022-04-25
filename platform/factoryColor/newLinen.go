package factoryColor

import "image/color"

func NewLinen() color.RGBA {
	return color.RGBA{R: 0xfa, G: 0xf0, B: 0xe6, A: 0xff} // rgb(250, 240, 230)
}
