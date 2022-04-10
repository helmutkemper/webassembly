package canvas

import (
	"image/color"
	"syscall/js"
)

// todo: documentation

func (el *Canvas) GetImageDataPixelByCoordinate(data interface{}, x, y, width int) color.RGBA {
	width *= 4
	index := y*width + x*4

	if data.(js.Value).Index(index+0).IsUndefined() == true ||
		data.(js.Value).Index(index+1).IsUndefined() == true ||
		data.(js.Value).Index(index+2).IsUndefined() == true ||
		data.(js.Value).Index(index+3).IsUndefined() == true {
		return color.RGBA{}
	}

	return color.RGBA{
		R: uint8(data.(js.Value).Index(index + 0).Int()),
		G: uint8(data.(js.Value).Index(index + 1).Int()),
		B: uint8(data.(js.Value).Index(index + 2).Int()),
		A: uint8(data.(js.Value).Index(index + 3).Int()),
	}
}
