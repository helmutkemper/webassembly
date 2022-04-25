package factoryColor

import "image/color"

func NewLightsalmonTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xa0, B: 0x7a, A: 0x00} // rgb(255, 160, 122)
}
