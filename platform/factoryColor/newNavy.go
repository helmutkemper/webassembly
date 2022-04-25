package factoryColor

import "image/color"

func NewNavy() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x00, B: 0x80, A: 0xff} // rgb(0, 0, 128)
}
