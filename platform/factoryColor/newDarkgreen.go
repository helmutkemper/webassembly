package factoryColor

import "image/color"

func NewDarkgreen() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x64, B: 0x00, A: 0xff} // rgb(0, 100, 0)
}
