package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

// en: Draws an image, canvas, or video onto the canvas
//     img:     Specifies the image, canvas, or video element to use
//     sx:      Optional. The x coordinate where to start clipping
//     sy:      Optional. The y coordinate where to start clipping
//     swidth:  Optional. The width of the clipped image
//     sheight: Optional. The height of the clipped image
//     x:       The x coordinate where to place the image on the canvas
//     y:       The y coordinate where to place the image on the canvas
//     width:   Optional. The width of the image to use (stretch or reduce the image)
//     height:  Optional. The height of the image to use (stretch or reduce the image)
//
//     The drawImage() method draws an image, canvas, or video onto the canvas.
//     The drawImage() method can also draw parts of an image, and/or increase/reduce the image size.
//
//     Position the image on the canvas:
//     JavaScript syntax: context.drawImage(img, x, y);
//
//     Position the image on the canvas, and specify width and height of the image:
//     JavaScript syntax: context.drawImage(img, x, y, width, height);
//
//     Clip the image and position the clipped part on the canvas:
//     JavaScript syntax: context.drawImage(img, sx, sy, swidth, sheight, x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var img = document.getElementById("scream");
//     ctx.drawImage(img, 10, 10);
func (el *Canvas) DrawImage(value DrawImage) {

	if value.SX != 0 || value.SY != 0 || value.SWidth != 0 || value.SHeight != 0 {
		//context.drawImage(img, sx, sy, swidth, sheight, x, y, width, height);
		el.selfContext.Call("drawImage", value.Image, value.SX, value.SY, value.SWidth, value.SHeight, value.X, value.Y, value.Width, value.Height)
	} else if value.Width != 0 || value.Height != 0 {
		//context.drawImage(img, x, y, width, height);
		el.selfContext.Call("drawImage", value.Image, value.X, value.Y, value.Width, value.Height)
	} else {
		//context.drawImage(img, x, y);
		el.selfContext.Call("drawImage", value.Image, value.X, value.Y)
	}
}
