package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Creates a linear gradient (to use on canvas content)
//     x0: The x-coordinate of the start point of the gradient
//     y0: The y-coordinate of the start point of the gradient
//     x1: The x-coordinate of the end point of the gradient
//     y1: The y-coordinate of the end point of the gradient
//
//     The createLinearGradient() method creates a linear gradient object.
//     The gradient can be used to fill rectangles, circles, lines, text, etc.
//     Tip: Use this object as the value to the strokeStyle or fillStyle properties.
//     Tip: Use the addColorStop() method to specify different colors, and where to position the colors in the gradient object.
//     JavaScript syntax:	context.createLinearGradient(x0, y0, x1, y1);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var grd = ctx.createLinearGradient(0, 0, 170, 0);
//     grd.addColorStop(0, "black");
//     grd.addColorStop(1, "white");
//     ctx.fillStyle = grd;
//     ctx.fillRect(20, 20, 150, 100);
func (el *Canvas) CreateLinearGradient(x0, y0, x1, y1 iotmaker_types.Coordinate) {
	el.SelfContext.Call("createLinearGradient", x0, y0, x1, y1)
}
