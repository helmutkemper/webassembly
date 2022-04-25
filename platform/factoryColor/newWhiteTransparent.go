package factoryColor

import "image/color"

func NewWhiteTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x00} // rgb(255, 255, 255)
}
