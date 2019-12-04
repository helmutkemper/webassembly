package canvas

import (
	"image/color"
)

// en: Returns an ImageData object that copies the pixel data for the specified rectangle on a canvas
//     x:      The x coordinate (in pixels) of the upper-left corner to start copy from
//     y:      The y coordinate (in pixels) of the upper-left corner to start copy from
//     width:  The width of the rectangular area you will copy
//     height: The height of the rectangular area you will copy
//     return: map[x(int)][y(int)]alpha channel(uint8)
//             Note: return x and y are NOT relative to the coordinate (0,0) on the image, are relative to the
//                   coordinate (0,0) on the canvas
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
func (el *Canvas) GetImageDataAlphaChannelOnly(x, y, width, height int) map[int]map[int]uint8 {

	dataInterface := el.SelfContext.Call("getImageData", x, y, width, height)
	dataJs := dataInterface.Get("data")

	ret := make(map[int]map[int]uint8)

	var rgbaLength int = 4

	var tmp uint8
	var i int = 0
	var xp int
	var yp int
	for yp = 0; yp != height; yp += 1 {
		for xp = 0; xp != width; xp += 1 {

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			tmp = uint8(dataJs.Index(i + 3).Int())

			i += rgbaLength

			if tmp == 0 {
				continue
			}

			if len(ret[x+xp]) == 0 {
				ret[x+xp] = make(map[int]uint8)
			}

			ret[x+xp][y+yp] = tmp
		}
	}

	return ret
}
