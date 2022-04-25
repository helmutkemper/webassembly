package factoryColor

import "image/color"

func NewLightcoralTransparent() color.RGBA {
	return color.RGBA{R: 0xf0, G: 0x80, B: 0x80, A: 0x00} // rgb(240, 128, 128)
}
