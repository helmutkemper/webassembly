package factoryColor

import "image/color"

func NewPeru() color.RGBA {
	return color.RGBA{R: 0xcd, G: 0x85, B: 0x3f, A: 0xff} // rgb(205, 133, 63)
}
