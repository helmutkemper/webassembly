package canvas

// PutImageDataJsValue
//
// English:
//
//  PutImageDataJsValue() method puts the image data (from a specified ImageData object) back onto the
//  canvas.
//
//   Input:
//     imgData:     specifies the ImageData object to put back onto the canvas;
//     x:           the x-coordinate, in pixels, of the upper-left corner of where to place the image
//                  on the canvas;
//     y:           the y-coordinate, in pixels, of the upper-left corner of where to place the image
//                  on the canvas;
//     dirtyX:      optional. The x-coordinate, in pixels, of the upper-left corner of where to start
//                  drawing the image. Default 0;
//     dirtyY:      optional. The y-coordinate, in pixels, of the upper-left corner of where to start
//                  drawing the image. Default 0;
//     dirtyWidth:  optional. The width to use when drawing the image on the canvas;
//     Default:     the width of the extracted image;
//     dirtyHeight: optional. The height to use when drawing the image on the canvas;
//     Default:     the height of the extracted image.
//
//   Tip:
//     * Read about the getImageDataJsValue() method that copies the pixel data for a specified
//       rectangle on a canvas;
//     * Read about the createImageJsValue() method that creates a new, blank ImageData object.
func (el *Canvas) PutImageDataJsValue(data interface{}, values ...int) {
	if len(values) == 2 {
		x := values[0]
		y := values[1]
		el.SelfContext.Call("putImageData", data, x, y)
	} else if len(values) == 6 {
		x := values[0]
		y := values[1]
		dirtyX := values[2]
		dirtyY := values[3]
		dirtyWidth := values[4]
		dirtyHeight := values[5]
		el.SelfContext.Call("putImageData", data, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
	}
}
