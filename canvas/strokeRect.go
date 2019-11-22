package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Draws a rectangle (no fill)
//     x:      The x-coordinate of the upper-left corner of the rectangle
//     y:      The y-coordinate of the upper-left corner of the rectangle
//     width:  The width of the rectangle, in pixels
//     height: The height of the rectangle, in pixels
//
//     The strokeRect() method draws a rectangle (no fill). The default color of the stroke is black.
//     Tip: Use the strokeStyle property to set a color, gradient, or pattern to style the stroke.
//     JavaScript syntax: context.strokeRect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.strokeRect(20, 20, 150, 100);
func (el *Canvas) StrokeRect(x, y, width, height iotmaker_types.Coordinate) {
	el.SelfContext.Call("strokeRect", x, y, width, height)
}
