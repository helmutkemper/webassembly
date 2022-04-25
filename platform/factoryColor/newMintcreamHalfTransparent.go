package factoryColor

import "image/color"

func NewMintcreamHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xf5, G: 0xff, B: 0xfa, A: 0x80} // rgb(245, 255, 250)
}
