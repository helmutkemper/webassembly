package factoryColor

import "image/color"

func NewIvoryTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xff, B: 0xf0, A: 0x00} // rgb(255, 255, 240)
}
