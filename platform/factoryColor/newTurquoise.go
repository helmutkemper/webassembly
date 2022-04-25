package factoryColor

import "image/color"

func NewTurquoise() color.RGBA {
	return color.RGBA{R: 0x40, G: 0xe0, B: 0xd0, A: 0xff} // rgb(64, 224, 208)
}
