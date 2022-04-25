package factoryColor

import "image/color"

func NewDimgrey() color.RGBA {
	return color.RGBA{R: 0x69, G: 0x69, B: 0x69, A: 0xff} // rgb(105, 105, 105)
}
