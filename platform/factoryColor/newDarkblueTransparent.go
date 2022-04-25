package factoryColor

import "image/color"

func NewDarkblueTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x00, B: 0x8b, A: 0x00} // rgb(0, 0, 139)
}
