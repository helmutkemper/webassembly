package canvas

// en: Returns the width of an ImageData object
//
//     The width property returns the width of an ImageData object, in pixels.
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//     JavaScript syntax: imgData.width;
func (el *Canvas) GetWidth() float64 {
	return el.SelfContext.Get("width").Float()
}
