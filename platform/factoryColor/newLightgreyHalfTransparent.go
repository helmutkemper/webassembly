package factoryColor

import "image/color"

func NewLightgreyHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xd3, G: 0xd3, B: 0xd3, A: 0x80} // rgb(211, 211, 211)
}
