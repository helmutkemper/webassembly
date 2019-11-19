package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

// en: Returns the width of an ImageData object
//
//     The width property returns the width of an ImageData object, in pixels.
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//     JavaScript syntax: imgData.width;
func (el *Canvas) Width() float64 {
	return el.selfContext.Get("width").Float()
}

// en: Returns the height of an ImageData object
//
//     The height property returns the height of an ImageData object, in pixels.
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//     JavaScript syntax: imgData.height;
func (el *Canvas) Height() float64 {
	return el.selfContext.Get("height").Float()
}

// en: Returns an object that contains image data of a specified ImageData object
//
//     JavaScript syntax: imageData.data;
//     The data property returns an object that contains image data of the specified ImageData object.
//     For every pixel in an ImageData object there are four pieces of information, the RGBA values:
//
//     R - The color red (from 0-255)
//     G - The color green (from 0-255)
//     B - The color blue (from 0-255)
//     A - The alpha channel (from 0-255; 0 is transparent and 255 is fully visible)
//
//     The color/alpha information is held in an array, and is stored in the data property of the ImageData object.
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
//     Tip: Look at createImageData(), getImageData(), and putImageData() to learn more about the ImageData object.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var imgData = ctx.createImageData(100, 100);
//     for (var i = 0; i < imgData.data.length; i += 4) {
//       imgData.data[i+0] = 255;
//       imgData.data[i+1] = 0;
//       imgData.data[i+2] = 0;
//       imgData.data[i+3] = 255;
//     }
//     ctx.putImageData(imgData, 10, 10);
func (el *Canvas) Data() js.Value {
	return el.selfContext.Get("data")
}

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
	el.selfContext.Call("createImageData", data)
}

// en: Returns an ImageData object that copies the pixel data for the specified rectangle on a canvas
//     x:      The x coordinate (in pixels) of the upper-left corner to start copy from
//     y:      The y coordinate (in pixels) of the upper-left corner to start copy from
//     width:  The width of the rectangular area you will copy
//     height: The height of the rectangular area you will copy
//
//     JavaScript syntax: context.getImageData(x, y, width, height);
//
//     The getImageData() method returns an ImageData object that copies the pixel data for the specified rectangle on a
//     canvas.
//     Note: The ImageData object is not a picture, it specifies a part (rectangle) on the canvas, and holds information
//     of every pixel inside that rectangle.
//
//     For every pixel in an ImageData object there are four pieces of information, the RGBA values:
//     R - The color red (from 0-255)
//     G - The color green (from 0-255)
//     B - The color blue (from 0-255)
//     A - The alpha channel (from 0-255; 0 is transparent and 255 is fully visible)
//
//     The color/alpha information is held in an array, and is stored in the data property of the ImageData object.
//     Tip: After you have manipulated the color/alpha information in the array, you can copy the image data back onto
//     the canvas with the putImageData() method.
//
//     Example:
//     The code for getting color/alpha information of the first pixel in the returned ImageData object:
//          red = imgData.data[0];
//          green = imgData.data[1];
//          blue = imgData.data[2];
//          alpha = imgData.data[3];
//
//     Tip: You can also use the getImageData() method to invert the color of every pixels of an image on the canvas.
//
//     Loop through all the pixels and change the color values using this formula:
//          red = 255-old_red;
//          green = 255-old_green;
//          blue = 255-old_blue;
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
func (el *Canvas) GetImageData(x, y, width, height float64) {
	el.selfContext.Call("getImageData", x, y, width, height)
}

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
func (el *Canvas) PutImageData(imgData js.Value, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight float64) {
	el.selfContext.Call("putImageData", imgData, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}
