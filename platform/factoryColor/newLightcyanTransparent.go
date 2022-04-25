package factoryColor

import "image/color"

func NewLightcyanTransparent() color.RGBA {
	return color.RGBA{R: 0xe0, G: 0xff, B: 0xff, A: 0x00} // rgb(224, 255, 255)
}
