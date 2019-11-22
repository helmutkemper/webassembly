package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Draws a "filled" rectangle
//     x:      The x-coordinate of the upper-left corner of the rectangle
//     y:      The y-coordinate of the upper-left corner of the rectangle
//     width:  The width of the rectangle, in pixels
//     height: The height of the rectangle, in pixels
//
//     The fillRect() method draws a "filled" rectangle. The default color of the fill is black.
//     Tip: Use the fillStyle property to set a color, gradient, or pattern used to fill the drawing.
//     JavaScript syntax: context.fillRect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillRect(20, 20, 150, 100);
func (el *Canvas) FillRect(x, y, width, height iotmaker_types.Pixel) {
	el.SelfContext.Call("fillRect", x, y, width, height)
}
