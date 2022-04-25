package factoryColor

import "image/color"

func NewFuchsia() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0xff} // rgb(255, 0, 255)
}
