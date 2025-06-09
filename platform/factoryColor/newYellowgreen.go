package factoryColor

import "image/color"

func NewYellowGreen() color.RGBA {
	return color.RGBA{R: 0x9a, G: 0xcd, B: 0x32, A: 0xff} // rgb(154, 205, 50)
}
