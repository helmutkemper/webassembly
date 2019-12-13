package canvas

import (
	"syscall/js"
)

// todo: documentation
func (el *Canvas) GetImageDataAlphaChannelByCoordinate(data interface{}, x, y, width int) uint8 {
	width *= 4
	index := y*width + x*4

	if data.(js.Value).Index(index+3) == js.Undefined() {
		return 0
	}

	return uint8(data.(js.Value).Index(index + 3).Int())
}
