package factoryColor

import "image/color"

func NewWhite() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} // rgb(255, 255, 255)
}
