package canvas

// en: Returns the height of an ImageData object
//
//     The height property returns the height of an ImageData object, in pixels.
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//     JavaScript syntax: imgData.height;
func (el *Canvas) GetHeight() float64 {
	return el.SelfContext.Get("height").Float()
}
