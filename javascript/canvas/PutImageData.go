package canvas

// PutImageData
// en: The putImageData() method puts the image data (from a specified ImageData
// object) back onto the canvas.
// Tip: Read about the getImageData() method that copies the pixel data for a
// specified rectangle on a canvas.
// Tip: Read about the createImageData() method that creates a new, blank ImageData
// object.
//     imgData: Specifies the ImageData object to put back onto the canvas
//     x: The x-coordinate, in pixels, of the upper-left corner of where to place the
//     image on the canvas
//     y: The y-coordinate, in pixels, of the upper-left corner of where to place the
//     image on the canvas
//     dirtyX: Optional. The x-coordinate, in pixels, of the upper-left corner of
//     where to start drawing the image. Default 0
//     dirtyY: Optional. The y-coordinate, in pixels, of the upper-left corner of
//     where to start drawing the image. Default 0
//     dirtyWidth: Optional. The width to use when drawing the image on the canvas.
//     Default: the width of the extracted image
//     dirtyHeight: Optional. The height to use when drawing the image on the canvas.
//     Default: the height of the extracted image
func (el *Canvas) PutImageData(imgData interface{}, values ...int) {
	if len(values) == 2 {
		x := values[0]
		y := values[1]
		el.SelfContext.Call("putImageData", imgData, x, y)
	} else if len(values) == 6 {
		x := values[0]
		y := values[1]
		dirtyX := values[2]
		dirtyY := values[3]
		dirtyWidth := values[4]
		dirtyHeight := values[5]
		el.SelfContext.Call("putImageData", imgData, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
	}
}
