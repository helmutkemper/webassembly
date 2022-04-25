package factoryColor

import "image/color"

func NewOrange() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xa5, B: 0x00, A: 0xff} // rgb(255, 165, 0)
}
