package factoryColor

import "image/color"

func NewBeigeTransparent() color.RGBA {
	return color.RGBA{R: 0xf5, G: 0xf5, B: 0xdc, A: 0x00} // rgb(245, 245, 220)
}
