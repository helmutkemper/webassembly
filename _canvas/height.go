package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Returns the height of an ImageData object
//
//     The height property returns the height of an ImageData object, in pixels.
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//     JavaScript syntax: imgData.height;
func (el *Canvas) Height() iotmaker_types.Coordinate {
	return el.SelfContext.Get("height").Float()
}
