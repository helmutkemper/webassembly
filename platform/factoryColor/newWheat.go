package factoryColor

import "image/color"

func NewWheat() color.RGBA {
	return color.RGBA{R: 0xf5, G: 0xde, B: 0xb3, A: 0xff} // rgb(245, 222, 179)
}
