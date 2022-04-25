package factoryColor

import "image/color"

func NewTomato() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x63, B: 0x47, A: 0xff} // rgb(255, 99, 71)
}
