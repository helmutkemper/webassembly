package factoryColor

import "image/color"

func NewTomatoHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x63, B: 0x47, A: 0x80} // rgb(255, 99, 71)
}
