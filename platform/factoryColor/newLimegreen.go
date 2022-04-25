package factoryColor

import "image/color"

func NewLimegreen() color.RGBA {
	return color.RGBA{R: 0x32, G: 0xcd, B: 0x32, A: 0xff} // rgb(50, 205, 50)
}
