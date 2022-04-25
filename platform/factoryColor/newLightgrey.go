package factoryColor

import "image/color"

func NewLightgrey() color.RGBA {
	return color.RGBA{R: 0xd3, G: 0xd3, B: 0xd3, A: 0xff} // rgb(211, 211, 211)
}
