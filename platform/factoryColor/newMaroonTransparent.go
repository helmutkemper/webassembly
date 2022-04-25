package factoryColor

import "image/color"

func NewMaroonTransparent() color.RGBA {
	return color.RGBA{R: 0x80, G: 0x00, B: 0x00, A: 0x00} // rgb(128, 0, 0)
}
