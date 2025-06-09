package factoryColor

import "image/color"

func NewDarkMagenta() color.RGBA {
	return color.RGBA{R: 0x8b, G: 0x00, B: 0x8b, A: 0xff} // rgb(139, 0, 139)
}
