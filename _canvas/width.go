package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Returns the width of an ImageData object
//
//     The width property returns the width of an ImageData object, in pixels.
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//     JavaScript syntax: imgData.width;
func (el *Canvas) Width() iotmaker_types.Coordinate {
	return el.SelfContext.Get("width").Float()
}
