package canvas

import (
	"log"
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
func (el *Canvas) DrawImage(image interface{}, value ...int) {

	if len(value) == 8 {
		sx := value[0]
		sy := value[1]
		sWidth := value[2]
		sHeight := value[3]
		x := value[4]
		y := value[5]
		width := value[6]
		height := value[7]

		el.SelfContext.Call("drawImage", image, sx, sy, sWidth, sHeight, x, y, width, height)
	} else if len(value) == 4 {
		x := value[0]
		y := value[1]
		width := value[2]
		height := value[3]

		el.SelfContext.Call("drawImage", image, x, y, width, height)
	} else if len(value) == 2 {
		x := value[0]
		y := value[1]

		el.SelfContext.Call("drawImage", image.(js.Value), x, y)
	} else {
		log.Fatalf("canvas.drawImage must be canvas.drawImage(image, sx, sy, sWidth, sHeight, x, y, width, height), canvas.drawImage(image, x, y, width, height) or canvas.drawImage(image, x, y)")
	}
}
