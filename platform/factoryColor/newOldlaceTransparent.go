package factoryColor

import "image/color"

func NewOldlaceTransparent() color.RGBA {
	return color.RGBA{R: 0xfd, G: 0xf5, B: 0xe6, A: 0x00} // rgb(253, 245, 230)
}
