package factoryColor

import "image/color"

func NewLightgreenHalfTransparent() color.RGBA {
	return color.RGBA{R: 0x90, G: 0xee, B: 0x90, A: 0x80} // rgb(144, 238, 144)
}
