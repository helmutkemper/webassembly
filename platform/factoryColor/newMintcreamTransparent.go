package factoryColor

import "image/color"

func NewMintcreamTransparent() color.RGBA {
	return color.RGBA{R: 0xf5, G: 0xff, B: 0xfa, A: 0x00} // rgb(245, 255, 250)
}
