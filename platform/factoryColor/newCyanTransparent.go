package factoryColor

import "image/color"

func NewCyanTransparent() color.RGBA {
	return color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0x00} // rgb(0, 255, 255)
}
