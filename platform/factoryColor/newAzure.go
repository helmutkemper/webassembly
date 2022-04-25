package factoryColor

import "image/color"

func NewAzure() color.RGBA {
	return color.RGBA{R: 0xf0, G: 0xff, B: 0xff, A: 0xff} // rgb(240, 255, 255)
}
