package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Creates a rectangle
//     x:      The x-coordinate of the upper-left corner of the rectangle
//     y:      The y-coordinate of the upper-left corner of the rectangle
//     width:  The width of the rectangle, in pixels
//     height: The height of the rectangle, in pixels
//
//     The rect() method creates a rectangle.
//     Tip: Use the stroke() or the fill() method to actually draw the rectangle on the canvas.
//     JavaScript syntax: context.rect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.rect(20, 20, 150, 100);
//     ctx.stroke();
func (el *Canvas) Rect(x, y, width, height iotmaker_types.Coordinate) {
	el.SelfContext.Call("rect", x, y, width, height)
}
