package factoryColor

import "image/color"

func NewOlive() color.RGBA {
	return color.RGBA{R: 0x80, G: 0x80, B: 0x00, A: 0xff} // rgb(128, 128, 0)
}
