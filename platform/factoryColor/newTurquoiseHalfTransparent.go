package factoryColor

import "image/color"

func NewTurquoiseHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x40, G: 0xe0, B: 0xd0, A: 0x80} // rgb(64, 224, 208)
}
