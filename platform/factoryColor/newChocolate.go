package factoryColor

import "image/color"

func NewChocolate() color.RGBA {
	return color.RGBA{R: 0xd2, G: 0x69, B: 0x1e, A: 0xff} // rgb(210, 105, 30)
}
