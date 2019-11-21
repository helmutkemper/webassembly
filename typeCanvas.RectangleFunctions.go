package iotmaker_platform_webbrowser

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
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
func (el *Canvas) Rect(x, y, width, height iotmaker_types.Pixel) {
	el.SelfContext.Call("rect", x, y, width, height)
}

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
func (el *Canvas) fillRect(x, y, width, height iotmaker_types.Pixel) {
	el.SelfContext.Call("fillRect", x, y, width, height)
}

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
func (el *Canvas) StrokeRect(x, y, width, height iotmaker_types.Pixel) {
	el.SelfContext.Call("strokeRect", x, y, width, height)
}

// en: Clears the specified pixels within a given rectangle
//     x:      The x-coordinate of the upper-left corner of the rectangle to clear
//     y:      The y-coordinate of the upper-left corner of the rectangle to clear
//     width:  The width of the rectangle to clear, in pixels
//     height: The height of the rectangle to clear, in pixels
//
//     The clearRect() method clears the specified pixels within a given rectangle.
//     JavaScript syntax: context.clearRect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 300, 150);
//     ctx.clearRect(20, 20, 100, 50);
func (el *Canvas) ClearRect(x, y, width, height iotmaker_types.Pixel) {
	el.SelfContext.Call("clearRect", x, y, width, height)
}
