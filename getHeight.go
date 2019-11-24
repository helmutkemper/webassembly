package iotmaker_platform_webbrowser

// en: Returns the height of an ImageData object
//
//     The height property returns the height of an ImageData object, in pixels.
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//     JavaScript syntax: imgData.height;
func (el *Canvas) GetHeight() int {
	return el.SelfContext.Get("height").Int()
}
