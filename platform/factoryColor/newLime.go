package factoryColor

import "image/color"

func NewLime() color.RGBA {
	return color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff} // rgb(0, 255, 0)
}
