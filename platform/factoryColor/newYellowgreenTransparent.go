package factoryColor

import "image/color"

func NewYellowgreenTransparent() color.RGBA {
	return color.RGBA{R: 0x9a, G: 0xcd, B: 0x32, A: 0x00} // rgb(154, 205, 50)
}
