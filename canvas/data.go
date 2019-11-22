package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

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
	return el.SelfContext.Get("data")
}
