package factoryColor

import "image/color"

func NewPlum() color.RGBA {
	return color.RGBA{R: 0xdd, G: 0xa0, B: 0xdd, A: 0xff} // rgb(221, 160, 221)
}
