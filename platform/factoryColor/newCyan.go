package factoryColor

import "image/color"

func NewCyan() color.RGBA {
	return color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff} // rgb(0, 255, 255)
}
