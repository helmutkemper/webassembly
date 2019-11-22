package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Puts the image data (from a specified ImageData object) back onto the canvas
//     imgData:     Specifies the ImageData object to put back onto the canvas
//     x:           The x-coordinate, in pixels, of the upper-left corner of where to place the image on the canvas
//     y:           The y-coordinate, in pixels, of the upper-left corner of where to place the image on the canvas
//     dirtyX:      Optional. The x-coordinate, in pixels, of the upper-left corner of where to start drawing the image.
//                  Default 0
//     dirtyY:      Optional. The y-coordinate, in pixels, of the upper-left corner of where to start drawing the image.
//                  Default 0
//     dirtyWidth:  Optional. The width to use when drawing the image on the canvas. Default: the width of the extracted
//                  image
//     dirtyHeight: Optional. The height to use when drawing the image on the canvas. Default: the height of the
//                  extracted image
//
//     JavaScript syntax: context.putImageData(imgData, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight);
//
//     The putImageData() method puts the image data (from a specified ImageData object) back onto the canvas.
//     Tip: Read about the getImageData() method that copies the pixel data for a specified rectangle on a canvas.
//     Tip: Read about the createImageData() method that creates a new, blank ImageData object.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "red";
//     ctx.fillRect(10, 10, 50, 50);
//     function copy() {
//       var imgData = ctx.getImageData(10, 10, 50, 50);
//       ctx.putImageData(imgData, 10, 70);
//     }
//todo: fazer
func (el *Canvas) PutImageData(imgData js.Value, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight iotmaker_types.Pixel) {
	el.SelfContext.Call("putImageData", imgData, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}
