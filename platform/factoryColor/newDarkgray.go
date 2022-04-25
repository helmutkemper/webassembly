package factoryColor

import "image/color"

func NewDarkgray() color.RGBA {
	return color.RGBA{R: 0xa9, G: 0xa9, B: 0xa9, A: 0xff} // rgb(169, 169, 169)
}
