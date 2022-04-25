package factoryColor

import "image/color"

func NewAzureHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xf0, G: 0xff, B: 0xff, A: 0x80} // rgb(240, 255, 255)
}
