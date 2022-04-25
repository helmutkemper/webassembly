package factoryColor

import "image/color"

func NewBrown() color.RGBA {
	return color.RGBA{R: 0xa5, G: 0x2a, B: 0x2a, A: 0xff} // rgb(165, 42, 42)
}
