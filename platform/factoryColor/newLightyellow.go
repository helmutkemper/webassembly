package factoryColor

import "image/color"

func NewLightyellow() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xff, B: 0xe0, A: 0xff} // rgb(255, 255, 224)
}
