package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Creates a new, blank ImageData object
//     width:     The width of the new ImageData object, in pixels
//     height:    The height of the new ImageData object, in pixels
//     imageData: AnotherImageData object
//
//     There are two versions of the createImageData() method:
//     1. This creates a new ImageData object with the specified dimensions (in pixels):
//     JavaScript syntax: var imgData = context.createImageData(width, height);
//
//     2. This creates a new ImageData object with the same dimensions as the object specified by anotherImageData (this
//     does not copy the image data):
//     JavaScript syntax: var imgData = context.createImageData(imageData);
//
//     The createImageData() method creates a new, blank ImageData object. The new object's pixel values are transparent
//     black by default.
//     For every pixel in an ImageData object there are four pieces of information, the RGBA values:
//
//     R - The color red (from 0-255)
//     G - The color green (from 0-255)
//     B - The color blue (from 0-255)
//     A - The alpha channel (from 0-255; 0 is transparent and 255 is fully visible)
//
//     So, transparent black indicates: (0, 0, 0, 0).
//     The color/alpha information is held in an array, and since the array contains 4 pieces of information for every
//     pixel, the array's size is 4 times the size of the ImageData object: width*height*4. (An easier way to find the
//     size of the array, is to use ImageDataObject.data.length)
//
//     The array containing the color/alpha information is stored in the data property of the ImageData object.
//     Tip: After you have manipulated the color/alpha information in the array, you can copy the image data back onto
//     the canvas with the putImageData() method.
//
//     Examples:
//
//     The syntax for making the first pixel in the ImageData object red:
//          imgData = ctx.createImageData(100, 100);
//          imgData.data[0] = 255;
//          imgData.data[1] = 0;
//          imgData.data[2] = 0;
//          imgData.data[3] = 255;
//
//     The syntax for making the second pixel in the ImageData object green:
//          imgData = ctx.createImageData(100, 100);
//          imgData.data[4] = 0;
//          imgData.data[5] = 255;
//          imgData.data[6] = 0;
//          imgData.data[7] = 255;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var imgData = ctx.createImageData(100, 100);
//     for (var i = 0; i < imgData.data.length; i += 4)
//     {
//       imgData.data[i+0] = 255;
//       imgData.data[i+1] = 0;
//       imgData.data[i+2] = 0;
//       imgData.data[i+3] = 255;
//     }
//     ctx.putImageData(imgData, 10, 10);
//todo: fazer
func (el *Canvas) CreateImageData(data js.Value) {
	el.SelfContext.Call("createImageData", data)
}
